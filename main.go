package main

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/big"
	"net/url"
	"os"
	"os/exec"
	"regexp"
	"slices"
	"sort"
	"strings"

	"gitee.com/masx200/github-mcp-server/internal/ghmcp"
	"gitee.com/masx200/github-mcp-server/pkg/github"
	"gitee.com/masx200/github-mcp-server/pkg/raw"
	"gitee.com/masx200/github-mcp-server/pkg/toolsets"
	"gitee.com/masx200/github-mcp-server/pkg/translations"
	"gitee.com/masx200/github-mcp-server/streamablehttp"
	gogithub "github.com/google/go-github/v73/github"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/shurcooL/githubv4"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// These variables are set by the build process using ldflags.
var version = "version"
var commit = "commit"
var date = "date"

var (
	rootCmd = &cobra.Command{
		Use:     "server",
		Short:   "GitHub MCP Server",
		Long:    `A GitHub MCP server that handles various tools and resources.`,
		Version: fmt.Sprintf("Version: %s\nCommit: %s\nBuild Date: %s", version, commit, date),
	}

	stdioCmd = &cobra.Command{
		Use:   "stdio",
		Short: "Start stdio server",
		Long:  `Start a server that communicates via standard input/output streams using JSON-RPC messages.`,
		RunE: func(_ *cobra.Command, _ []string) error {
			token := viper.GetString("personal_access_token")
			if token == "" {
				return errors.New("GITHUB_PERSONAL_ACCESS_TOKEN not set")
			}

			// If you're wondering why we're not using viper.GetStringSlice("toolsets"),
			// it's because viper doesn't handle comma-separated values correctly for env
			// vars when using GetStringSlice.
			// https://github.com/spf13/viper/issues/380
			var enabledToolsets []string
			if err := viper.UnmarshalKey("toolsets", &enabledToolsets); err != nil {
				return fmt.Errorf("failed to unmarshal toolsets: %w", err)
			}

			stdioServerConfig := ghmcp.StdioServerConfig{
				Pretty:               viper.GetBool("pretty"),
				Version:              version,
				Host:                 viper.GetString("host"),
				Token:                token,
				EnabledToolsets:      enabledToolsets,
				DynamicToolsets:      viper.GetBool("dynamic_toolsets"),
				ReadOnly:             viper.GetBool("read-only"),
				ExportTranslations:   viper.GetBool("export-translations"),
				EnableCommandLogging: viper.GetBool("enable-command-logging"),
				LogFilePath:          viper.GetString("log-file"),
			}
			return ghmcp.RunStdioServer(stdioServerConfig)
		},
	}
)

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.SetGlobalNormalizationFunc(wordSepNormalizeFunc)

	rootCmd.SetVersionTemplate("{{.Short}}\n{{.Version}}\n")

	// Add global flags that will be shared by all commands
	rootCmd.PersistentFlags().StringSlice("toolsets", github.DefaultTools, "An optional comma separated list of groups of tools to allow, defaults to enabling all")
	rootCmd.PersistentFlags().Bool("dynamic-toolsets", false, "Enable dynamic toolsets")
	rootCmd.PersistentFlags().Bool("read-only", false, "Restrict the server to read-only operations")
	rootCmd.PersistentFlags().String("log-file", "", "Path to log file")
	rootCmd.PersistentFlags().Bool("enable-command-logging", false, "When enabled, the server will log all command requests and responses to the log file")
	rootCmd.PersistentFlags().Bool("export-translations", false, "Save translations to a JSON file")
	rootCmd.PersistentFlags().String("gh-host", "", "Specify the GitHub hostname (for GitHub Enterprise etc.)")

	//when transport is streamablehttp
	rootCmd.PersistentFlags().String("address", ":8080", "Address to listen on(when transport is streamablehttp)(default: \":8080\")")
	// Bind flag to viper

	//when transport is streamablehttp
	_ = viper.BindPFlag("address", rootCmd.PersistentFlags().Lookup("address"))
	_ = viper.BindPFlag("toolsets", rootCmd.PersistentFlags().Lookup("toolsets"))
	_ = viper.BindPFlag("dynamic_toolsets", rootCmd.PersistentFlags().Lookup("dynamic-toolsets"))
	_ = viper.BindPFlag("read-only", rootCmd.PersistentFlags().Lookup("read-only"))
	_ = viper.BindPFlag("log-file", rootCmd.PersistentFlags().Lookup("log-file"))
	_ = viper.BindPFlag("enable-command-logging", rootCmd.PersistentFlags().Lookup("enable-command-logging"))
	_ = viper.BindPFlag("export-translations", rootCmd.PersistentFlags().Lookup("export-translations"))
	_ = viper.BindPFlag("host", rootCmd.PersistentFlags().Lookup("gh-host"))

	// Add subcommands
	rootCmd.AddCommand(stdioCmd)
}

func initConfig() {
	// Initialize Viper configuration
	viper.SetEnvPrefix("github")
	viper.AutomaticEnv()

}

type PropertyItem struct {
	Type                 string              `json:"type"`
	Properties           map[string]Property `json:"properties,omitempty"`
	Required             []string            `json:"required,omitempty"`
	AdditionalProperties bool                `json:"additionalProperties,omitempty"`
}
type Property struct {
	Type        string        `json:"type"`
	Description string        `json:"description"`
	Enum        []string      `json:"enum,omitempty"`
	Minimum     *float64      `json:"minimum,omitempty"`
	Maximum     *float64      `json:"maximum,omitempty"`
	Items       *PropertyItem `json:"items,omitempty"`
}
type InputSchema struct {
	Type                 string              `json:"type"`
	Properties           map[string]Property `json:"properties"`
	Required             []string            `json:"required"`
	AdditionalProperties bool                `json:"additionalProperties"`
	Schema               string              `json:"$schema"`
}
type Tool struct {
	Name        string      `json:"name"`
	Description string      `json:"description"`
	InputSchema InputSchema `json:"inputSchema"`
}
type Result struct {
	Tools []Tool `json:"tools"`
}
type SchemaResponse struct {
	Result  Result `json:"result"`
	JSONRPC string `json:"jsonrpc"`
	ID      int    `json:"id"`
}

func main() {

	// Add schema command
	var schemaCmd = &cobra.Command{
		Use:   "schema",
		Short: "Fetch schema from MCP server",
		Long:  "Fetches the tools schema from the MCP server specified by --stdio-server-cmd",
		RunE: func(cmd *cobra.Command, _ []string) error {
			serverCmd, _ := cmd.Flags().GetString("stdio-server-cmd")
			if serverCmd == "" {
				return fmt.Errorf("--stdio-server-cmd is required")
			}

			// Build the JSON-RPC request for tools/list
			jsonRequest, err := buildJSONRPCRequest("tools/list", "", nil)
			if err != nil {
				return fmt.Errorf("failed to build JSON-RPC request: %w", err)
			}

			// Execute the server command and pass the JSON-RPC request
			response, err := executeServerCommand(serverCmd, jsonRequest)
			if err != nil {
				return fmt.Errorf("error executing server command: %w", err)
			}

			// Output the response
			fmt.Println(response)
			return nil
		},

		PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
			// Skip validation for help and completion commands
			if cmd.Name() == "help" || cmd.Name() == "completion" {
				return nil
			}

			// Check if the required global flag is provided
			serverCmd, _ := cmd.Flags().GetString("stdio-server-cmd")
			if serverCmd == "" {
				return fmt.Errorf("--stdio-server-cmd is required")
			}
			return nil
		},
	}

	// Create the tools command
	var toolsCmd = &cobra.Command{
		Use:   "tools",
		Short: "Access available tools",
		Long:  "Contains all dynamically generated tool commands from the schema",
		PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
			// Skip validation for help and completion commands
			if cmd.Name() == "help" || cmd.Name() == "completion" {
				return nil
			}

			// Check if the required global flag is provided
			serverCmd, _ := cmd.Flags().GetString("stdio-server-cmd")
			if serverCmd == "" {
				return fmt.Errorf("--stdio-server-cmd is required")
			}
			return nil
		},
	}
	var generateDocsCmd = &cobra.Command{
		Use:   "generate-docs",
		Short: "Generate documentation for tools and toolsets",
		Long:  `Generate the automated sections of README.md and docs/remote-server.md with current tool and toolset information.`,
		RunE: func(_ *cobra.Command, _ []string) error {
			return generateAllDocs()
		},
	}
	_ = schemaCmd.MarkPersistentFlagRequired("stdio-server-cmd")
	_ = toolsCmd.MarkPersistentFlagRequired("stdio-server-cmd")
	rootCmd.AddCommand(generateDocsCmd)
	rootCmd.AddCommand(schemaCmd)
	rootCmd.AddCommand(toolsCmd)

	rootCmd.PersistentFlags().String("stdio-server-cmd", "", "Shell command to invoke MCP server via stdio (required)")

	// Add global flag for pretty printing
	rootCmd.PersistentFlags().Bool("pretty", true, "Pretty print MCP response (only for JSON or JSONL responses)")

	_ = rootCmd.ParseFlags(os.Args[1:])
	_ = viper.BindPFlag("pretty", rootCmd.PersistentFlags().Lookup("pretty"))
	prettyPrint, err := rootCmd.Flags().GetBool("pretty")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error getting pretty flag: %v\n", err)
		os.Exit(1)
	}
	// Get server command
	serverCmd, err := rootCmd.Flags().GetString("stdio-server-cmd")
	if err == nil && serverCmd != "" {
		// Fetch schema from server
		jsonRequest, err := buildJSONRPCRequest("tools/list", "", nil)
		if err == nil {
			response, err := executeServerCommand(serverCmd, jsonRequest)
			if err == nil {
				// Parse the schema response
				var schemaResp SchemaResponse
				if err := json.Unmarshal([]byte(response), &schemaResp); err == nil {
					// Add all the generated commands as subcommands of tools
					for _, tool := range schemaResp.Result.Tools {
						addCommandFromTool(toolsCmd, &tool, prettyPrint)
					}
				}
			}
		}
	}
	var httpcmd = streamablehttp.HttpCmdfactory(version)
	rootCmd.AddCommand(httpcmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func buildArgumentsMap(cmd *cobra.Command, tool *Tool) (map[string]interface{}, error) {
	arguments := make(map[string]interface{})

	for name, prop := range tool.InputSchema.Properties {
		switch prop.Type {
		case "string":
			if value, _ := cmd.Flags().GetString(name); value != "" {
				arguments[name] = value
			}
		case "number":
			if value, _ := cmd.Flags().GetFloat64(name); value != 0 {
				arguments[name] = value
			}
		case "integer":
			if value, _ := cmd.Flags().GetInt64(name); value != 0 {
				arguments[name] = value
			}
		case "boolean":
			// For boolean, we need to check if it was explicitly set
			if cmd.Flags().Changed(name) {
				value, _ := cmd.Flags().GetBool(name)
				arguments[name] = value
			}
		case "array":
			if prop.Items != nil {
				switch prop.Items.Type {
				case "string":
					if values, _ := cmd.Flags().GetStringSlice(name); len(values) > 0 {
						arguments[name] = values
					}
				case "object":
					if jsonStr, _ := cmd.Flags().GetString(name + "-json"); jsonStr != "" {
						var jsonArray []interface{}
						if err := json.Unmarshal([]byte(jsonStr), &jsonArray); err != nil {
							return nil, fmt.Errorf("error parsing JSON for %s: %w", name, err)
						}
						arguments[name] = jsonArray
					}
				}
			}
		}
	}

	return arguments, nil
}
func addCommandFromTool(toolsCmd *cobra.Command, tool *Tool, prettyPrint bool) {
	// Create command from tool
	cmd := &cobra.Command{
		Use:   tool.Name,
		Short: tool.Description,
		Run: func(cmd *cobra.Command, _ []string) {
			// Build a map of arguments from flags
			arguments, err := buildArgumentsMap(cmd, tool)
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "failed to build arguments map: %v\n", err)
				return
			}

			jsonData, err := buildJSONRPCRequest("tools/call", tool.Name, arguments)
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "failed to build JSONRPC request: %v\n", err)
				return
			}

			// Execute the server command
			serverCmd, err := cmd.Flags().GetString("stdio-server-cmd")
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "failed to get stdio-server-cmd: %v\n", err)
				return
			}
			response, err := executeServerCommand(serverCmd, jsonData)
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "error executing server command: %v\n", err)
				return
			}
			if err := printResponse(response, prettyPrint); err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "error printing response: %v\n", err)
				return
			}
		},
	}

	// Initialize viper for this command
	viperInit := func() {
		viper.Reset()
		viper.AutomaticEnv()
		viper.SetEnvPrefix(strings.ToUpper(tool.Name))
		viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	}

	// We'll call the init function directly instead of with cobra.OnInitialize
	// to avoid conflicts between commands
	viperInit()

	// Add flags based on schema properties
	for name, prop := range tool.InputSchema.Properties {
		isRequired := slices.Contains(tool.InputSchema.Required, name)

		// Enhance description to indicate if parameter is optional
		description := prop.Description
		if !isRequired {
			description += " (optional)"
		}

		switch prop.Type {
		case "string":
			cmd.Flags().String(name, "", description)
			if len(prop.Enum) > 0 {
				// Add validation in PreRun for enum values
				cmd.PreRunE = func(cmd *cobra.Command, _ []string) error {
					for flagName, property := range tool.InputSchema.Properties {
						if len(property.Enum) > 0 {
							value, _ := cmd.Flags().GetString(flagName)
							if value != "" && !slices.Contains(property.Enum, value) {
								return fmt.Errorf("%s must be one of: %s", flagName, strings.Join(property.Enum, ", "))
							}
						}
					}
					return nil
				}
			}
		case "number":
			cmd.Flags().Float64(name, 0, description)
		case "integer":
			cmd.Flags().Int64(name, 0, description)
		case "boolean":
			cmd.Flags().Bool(name, false, description)
		case "array":
			if prop.Items != nil {
				switch prop.Items.Type {
				case "string":
					cmd.Flags().StringSlice(name, []string{}, description)
				case "object":
					cmd.Flags().String(name+"-json", "", description+" (provide as JSON array)")
				}
			}
		}

		if isRequired {
			_ = cmd.MarkFlagRequired(name)
		}

		// Bind flag to viper
		_ = viper.BindPFlag(name, cmd.Flags().Lookup(name))
	}

	// Add command to root
	toolsCmd.AddCommand(cmd)
}

type Response struct {
	Result  ResponseResult `json:"result"`
	JSONRPC string         `json:"jsonrpc"`
	ID      int            `json:"id"`
}
type ResponseResult struct {
	Content []Content `json:"content"`
}
type Content struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

func printResponse(response string, prettyPrint bool) error {
	if !prettyPrint {
		fmt.Println(response)
		return nil
	}

	// Parse the JSON response
	var resp Response
	if err := json.Unmarshal([]byte(response), &resp); err != nil {
		return fmt.Errorf("failed to parse JSON: %w", err)
	}

	// Extract text from content items of type "text"
	for _, content := range resp.Result.Content {
		if content.Type == "text" {
			var textContentObj map[string]interface{}
			err := json.Unmarshal([]byte(content.Text), &textContentObj)

			if err == nil {
				prettyText, err := json.MarshalIndent(textContentObj, "", "  ")
				if err != nil {
					return fmt.Errorf("failed to pretty print text content: %w", err)
				}
				fmt.Println(string(prettyText))
				continue
			}

			// Fallback parsing as JSONL
			var textContentList []map[string]interface{}
			if err := json.Unmarshal([]byte(content.Text), &textContentList); err != nil {
				return fmt.Errorf("failed to parse text content as a list: %w", err)
			}
			prettyText, err := json.MarshalIndent(textContentList, "", "  ")
			if err != nil {
				return fmt.Errorf("failed to pretty print array content: %w", err)
			}
			fmt.Println(string(prettyText))
		}
	}

	// If no text content found, print the original response
	if len(resp.Result.Content) == 0 {
		fmt.Println(response)
	}

	return nil
}

func wordSepNormalizeFunc(_ *pflag.FlagSet, name string) pflag.NormalizedName {
	from := []string{"_"}
	to := "-"
	for _, sep := range from {
		name = strings.ReplaceAll(name, sep, to)
	}
	return pflag.NormalizedName(name)
}

func generateAllDocs() error {
	if err := generateReadmeDocs("README.md"); err != nil {
		return fmt.Errorf("failed to generate README docs: %w", err)
	}

	if err := generateRemoteServerDocs("docs/remote-server.md"); err != nil {
		return fmt.Errorf("failed to generate remote-server docs: %w", err)
	}

	return nil
}

func generateReadmeDocs(readmePath string) error {
	// Create translation helper
	t, _ := translations.TranslationHelper()

	// Create toolset group with mock clients
	tsg := github.DefaultToolsetGroup(false, mockGetClient, mockGetGQLClient, mockGetRawClient, t)

	// Generate toolsets documentation
	toolsetsDoc := generateToolsetsDoc(tsg)

	// Generate tools documentation
	toolsDoc := generateToolsDoc(tsg)

	// Read the current README.md
	// #nosec G304 - readmePath is controlled by command line flag, not user input
	content, err := os.ReadFile(readmePath)
	if err != nil {
		return fmt.Errorf("failed to read README.md: %w", err)
	}

	// Replace toolsets section
	updatedContent := replaceSection(string(content), "START AUTOMATED TOOLSETS", "END AUTOMATED TOOLSETS", toolsetsDoc)

	// Replace tools section
	updatedContent = replaceSection(updatedContent, "START AUTOMATED TOOLS", "END AUTOMATED TOOLS", toolsDoc)

	// Write back to file
	err = os.WriteFile(readmePath, []byte(updatedContent), 0600)
	if err != nil {
		return fmt.Errorf("failed to write README.md: %w", err)
	}

	fmt.Println("Successfully updated README.md with automated documentation")
	return nil
}
func mockGetRawClient(_ context.Context) (*raw.Client, error) {
	return nil, nil
}
func generateRemoteServerDocs(docsPath string) error {
	content, err := os.ReadFile(docsPath) //#nosec G304
	if err != nil {
		return fmt.Errorf("failed to read docs file: %w", err)
	}

	toolsetsDoc := generateRemoteToolsetsDoc()

	// Replace content between markers
	startMarker := "<!-- START AUTOMATED TOOLSETS -->"
	endMarker := "<!-- END AUTOMATED TOOLSETS -->"

	contentStr := string(content)
	startIndex := strings.Index(contentStr, startMarker)
	endIndex := strings.Index(contentStr, endMarker)

	if startIndex == -1 || endIndex == -1 {
		return fmt.Errorf("automation markers not found in %s", docsPath)
	}

	newContent := contentStr[:startIndex] + startMarker + "\n" + toolsetsDoc + "\n" + endMarker + contentStr[endIndex+len(endMarker):]

	return os.WriteFile(docsPath, []byte(newContent), 0600) //#nosec G306
}

type RequestParams struct {
	Name      string                 `json:"name"`
	Arguments map[string]interface{} `json:"arguments"`
}

type JSONRPCRequest struct {
	JSONRPC string        `json:"jsonrpc"`
	ID      int           `json:"id"`
	Method  string        `json:"method"`
	Params  RequestParams `json:"params"`
}

func buildJSONRPCRequest(method, toolName string, arguments map[string]interface{}) (string, error) {
	id, err := rand.Int(rand.Reader, big.NewInt(10000))
	if err != nil {
		return "", fmt.Errorf("failed to generate random ID: %w", err)
	}
	request := JSONRPCRequest{
		JSONRPC: "2.0",
		ID:      int(id.Int64()), // Random ID between 0 and 9999
		Method:  method,
		Params: RequestParams{
			Name:      toolName,
			Arguments: arguments,
		},
	}
	jsonData, err := json.Marshal(request)
	if err != nil {
		return "", fmt.Errorf("failed to marshal JSON request: %w", err)
	}
	return string(jsonData), nil
}

func executeServerCommand(cmdStr, jsonRequest string) (string, error) {
	// Split the command string into command and arguments
	cmdParts := strings.Fields(cmdStr)
	if len(cmdParts) == 0 {
		return "", fmt.Errorf("empty command")
	}

	cmd := exec.Command(cmdParts[0], cmdParts[1:]...) //nolint:gosec //mcpcurl is a test command that needs to execute arbitrary shell commands

	// Setup stdin pipe
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return "", fmt.Errorf("failed to create stdin pipe: %w", err)
	}

	// Setup stdout and stderr pipes
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	// Start the command
	if err := cmd.Start(); err != nil {
		return "", fmt.Errorf("failed to start command: %w", err)
	}

	// Write the JSON request to stdin
	if _, err := io.WriteString(stdin, jsonRequest+"\n"); err != nil {
		return "", fmt.Errorf("failed to write to stdin: %w", err)
	}
	_ = stdin.Close()

	// Wait for the command to complete
	if err := cmd.Wait(); err != nil {
		return "", fmt.Errorf("command failed: %w, stderr: %s", err, stderr.String())
	}

	return stdout.String(), nil
}
func mockGetClient(_ context.Context) (*gogithub.Client, error) {
	return gogithub.NewClient(nil), nil
}
func mockGetGQLClient(_ context.Context) (*githubv4.Client, error) {
	return githubv4.NewClient(nil), nil
}

func generateToolsetsDoc(tsg *toolsets.ToolsetGroup) string {
	var lines []string

	// Add table header and separator
	lines = append(lines, "| Toolset                 | Description                                                   |")
	lines = append(lines, "| ----------------------- | ------------------------------------------------------------- |")

	// Add the context toolset row (handled separately in README)
	lines = append(lines, "| `context`               | **Strongly recommended**: Tools that provide context about the current user and GitHub context you are operating in |")

	// Get all toolsets except context (which is handled separately above)
	var toolsetNames []string
	for name := range tsg.Toolsets {
		if name != "context" && name != "dynamic" { // Skip context and dynamic toolsets as they're handled separately
			toolsetNames = append(toolsetNames, name)
		}
	}

	// Sort toolset names for consistent output
	sort.Strings(toolsetNames)

	for _, name := range toolsetNames {
		toolset := tsg.Toolsets[name]
		lines = append(lines, fmt.Sprintf("| `%s` | %s |", name, toolset.Description))
	}

	return strings.Join(lines, "\n")
}

func generateToolsDoc(tsg *toolsets.ToolsetGroup) string {
	var sections []string

	// Get all toolset names and sort them alphabetically for deterministic order
	var toolsetNames []string
	for name := range tsg.Toolsets {
		if name != "dynamic" { // Skip dynamic toolset as it's handled separately
			toolsetNames = append(toolsetNames, name)
		}
	}
	sort.Strings(toolsetNames)

	for _, toolsetName := range toolsetNames {
		toolset := tsg.Toolsets[toolsetName]

		tools := toolset.GetAvailableTools()
		if len(tools) == 0 {
			continue
		}

		// Sort tools by name for deterministic order
		sort.Slice(tools, func(i, j int) bool {
			return tools[i].Tool.Name < tools[j].Tool.Name
		})

		// Generate section header - capitalize first letter and replace underscores
		sectionName := formatToolsetName(toolsetName)

		var toolDocs []string
		for _, serverTool := range tools {
			toolDoc := generateToolDoc(serverTool.Tool)
			toolDocs = append(toolDocs, toolDoc)
		}

		if len(toolDocs) > 0 {
			section := fmt.Sprintf("<details>\n\n<summary>%s</summary>\n\n%s\n\n</details>",
				sectionName, strings.Join(toolDocs, "\n\n"))
			sections = append(sections, section)
		}
	}

	return strings.Join(sections, "\n\n")
}

func generateRemoteToolsetsDoc() string {
	var buf strings.Builder

	// Create translation helper
	t, _ := translations.TranslationHelper()

	// Create toolset group with mock clients
	tsg := github.DefaultToolsetGroup(false, mockGetClient, mockGetGQLClient, mockGetRawClient, t)

	// Generate table header
	buf.WriteString("| Name           | Description                                      | API URL                                               | 1-Click Install (VS Code)                                                                                                                                                                                                 | Read-only Link                                                                                                 | 1-Click Read-only Install (VS Code)                                                                                                                                                                                                 |\n")
	buf.WriteString("|----------------|--------------------------------------------------|-------------------------------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|\n")

	// Get all toolsets
	toolsetNames := make([]string, 0, len(tsg.Toolsets))
	for name := range tsg.Toolsets {
		if name != "context" && name != "dynamic" { // Skip context and dynamic toolsets as they're handled separately
			toolsetNames = append(toolsetNames, name)
		}
	}
	sort.Strings(toolsetNames)

	// Add "all" toolset first (special case)
	buf.WriteString("| all            | All available GitHub MCP tools                    | https://api.githubcopilot.com/mcp/                    | [Install](https://insiders.vscode.dev/redirect/mcp/install?name=github&config=%7B%22type%22%3A%20%22http%22%2C%22url%22%3A%20%22https%3A%2F%2Fapi.githubcopilot.com%2Fmcp%2F%22%7D)                                      | [read-only](https://api.githubcopilot.com/mcp/readonly)                                                      | [Install read-only](https://insiders.vscode.dev/redirect/mcp/install?name=github&config=%7B%22type%22%3A%20%22http%22%2C%22url%22%3A%20%22https%3A%2F%2Fapi.githubcopilot.com%2Fmcp%2Freadonly%22%7D) |\n")

	// Add individual toolsets
	for _, name := range toolsetNames {
		toolset := tsg.Toolsets[name]

		formattedName := formatToolsetName(name)
		description := toolset.Description
		apiURL := fmt.Sprintf("https://api.githubcopilot.com/mcp/x/%s", name)
		readonlyURL := fmt.Sprintf("https://api.githubcopilot.com/mcp/x/%s/readonly", name)

		// Create install config JSON (URL encoded)
		installConfig := url.QueryEscape(fmt.Sprintf(`{"type": "http","url": "%s"}`, apiURL))
		readonlyConfig := url.QueryEscape(fmt.Sprintf(`{"type": "http","url": "%s"}`, readonlyURL))

		// Fix URL encoding to use %20 instead of + for spaces
		installConfig = strings.ReplaceAll(installConfig, "+", "%20")
		readonlyConfig = strings.ReplaceAll(readonlyConfig, "+", "%20")

		installLink := fmt.Sprintf("[Install](https://insiders.vscode.dev/redirect/mcp/install?name=gh-%s&config=%s)", name, installConfig)
		readonlyInstallLink := fmt.Sprintf("[Install read-only](https://insiders.vscode.dev/redirect/mcp/install?name=gh-%s&config=%s)", name, readonlyConfig)

		buf.WriteString(fmt.Sprintf("| %-14s | %-48s | %-53s | %-218s | %-110s | %-288s |\n",
			formattedName,
			description,
			apiURL,
			installLink,
			fmt.Sprintf("[read-only](%s)", readonlyURL),
			readonlyInstallLink,
		))
	}

	return buf.String()
}
func replaceSection(content, startMarker, endMarker, newContent string) string {
	startPattern := fmt.Sprintf(`<!-- %s -->`, regexp.QuoteMeta(startMarker))
	endPattern := fmt.Sprintf(`<!-- %s -->`, regexp.QuoteMeta(endMarker))

	re := regexp.MustCompile(fmt.Sprintf(`(?s)%s.*?%s`, startPattern, endPattern))

	replacement := fmt.Sprintf("<!-- %s -->\n%s\n<!-- %s -->", startMarker, newContent, endMarker)

	return re.ReplaceAllString(content, replacement)
}

func formatToolsetName(name string) string {
	switch name {
	case "pull_requests":
		return "Pull Requests"
	case "repos":
		return "Repositories"
	case "code_security":
		return "Code Security"
	case "secret_protection":
		return "Secret Protection"
	case "orgs":
		return "Organizations"
	default:
		// Fallback: capitalize first letter and replace underscores with spaces
		parts := strings.Split(name, "_")
		for i, part := range parts {
			if len(part) > 0 {
				parts[i] = strings.ToUpper(string(part[0])) + part[1:]
			}
		}
		return strings.Join(parts, " ")
	}
}

func generateToolDoc(tool mcp.Tool) string {
	var lines []string

	// Tool name only (using annotation name instead of verbose description)
	lines = append(lines, fmt.Sprintf("- **%s** - %s", tool.Name, tool.Annotations.Title))

	// Parameters
	schema := tool.InputSchema
	if len(schema.Properties) > 0 {
		// Get parameter names and sort them for deterministic order
		var paramNames []string
		for propName := range schema.Properties {
			paramNames = append(paramNames, propName)
		}
		sort.Strings(paramNames)

		for _, propName := range paramNames {
			prop := schema.Properties[propName]
			required := contains(schema.Required, propName)
			requiredStr := "optional"
			if required {
				requiredStr = "required"
			}

			// Get the type and description
			typeStr := "unknown"
			description := ""

			if propMap, ok := prop.(map[string]interface{}); ok {
				if typeVal, ok := propMap["type"].(string); ok {
					if typeVal == "array" {
						if items, ok := propMap["items"].(map[string]interface{}); ok {
							if itemType, ok := items["type"].(string); ok {
								typeStr = itemType + "[]"
							}
						} else {
							typeStr = "array"
						}
					} else {
						typeStr = typeVal
					}
				}

				if desc, ok := propMap["description"].(string); ok {
					description = desc
				}
			}

			paramLine := fmt.Sprintf("  - `%s`: %s (%s, %s)", propName, description, typeStr, requiredStr)
			lines = append(lines, paramLine)
		}
	} else {
		lines = append(lines, "  - No parameters required")
	}

	return strings.Join(lines, "\n")
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
