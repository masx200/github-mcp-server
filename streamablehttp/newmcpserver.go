package streamablehttp

import (
	"context"
	"fmt"
	"net/http"

	"gitee.com/masx200/github-mcp-server/pkg/errors"
	"gitee.com/masx200/github-mcp-server/pkg/github"
	"gitee.com/masx200/github-mcp-server/pkg/raw"
	gogithub "github.com/google/go-github/v73/github"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/shurcooL/githubv4"
)

func NewMCPServer(cfg MCPServerConfig) (*server.MCPServer, error) {
	apiHost, err := parseAPIHost(cfg.Host)
	if err != nil {
		return nil, fmt.Errorf("failed to parse API host: %w", err)
	}

	// Construct our REST client
	// WithAuthToken移到getClient中
	restClient := gogithub.NewClient(nil) /* .WithAuthToken(cfg.Token) */
	restClient.UserAgent = fmt.Sprintf("github-mcp-server/%s", cfg.Version)
	restClient.BaseURL = apiHost.baseRESTURL
	restClient.UploadURL = apiHost.uploadURL

	// Construct our GraphQL client
	// We're using NewEnterpriseClient here unconditionally as opposed to NewClient because we already
	// did the necessary API host parsing so that github.com will return the correct URL anyway.
	gqlHTTPClient := &http.Client{
		//bearerAuthTransport移到getGQLClient处理
		// Transport: &bearerAuthTransport{
		// 	transport: http.DefaultTransport,
		// 	token:     cfg.Token,
		// },
	} // We're going to wrap the Transport later in beforeInit
	gqlClient := NewEnterpriseClient(apiHost.graphqlURL.String(), gqlHTTPClient)

	// When a client send an initialize request, update the user agent to include the client info.
	beforeInit := func(_ context.Context, _ any, message *mcp.InitializeRequest) {
		userAgent := fmt.Sprintf(
			"github-mcp-server/%s (%s/%s)",
			cfg.Version,
			message.Params.ClientInfo.Name,
			message.Params.ClientInfo.Version,
		)

		restClient.UserAgent = userAgent

		gqlHTTPClient.Transport = &userAgentTransport{
			transport: gqlHTTPClient.Transport,
			agent:     userAgent,
		}
	}

	hooks := &server.Hooks{
		OnBeforeInitialize: []server.OnBeforeInitializeFunc{beforeInit},
		OnBeforeAny: []server.BeforeAnyHookFunc{
			func(ctx context.Context, _ any, _ mcp.MCPMethod, _ any) {
				// Ensure the context is cleared of any previous errors
				// as context isn't propagated through middleware
				errors.ContextWithGitHubErrors(ctx)
			},
		},
	}

	ghServer := github.NewServer(cfg.Version, server.WithHooks(hooks))

	enabledToolsets := cfg.EnabledToolsets
	if cfg.DynamicToolsets {
		// filter "all" from the enabled toolsets
		enabledToolsets = make([]string, 0, len(cfg.EnabledToolsets))
		for _, toolset := range cfg.EnabledToolsets {
			if toolset != "all" {
				enabledToolsets = append(enabledToolsets, toolset)
			}
		}
	}

	getClient := func(ctx context.Context) (*gogithub.Client, error) {
		token, err := cfg.Token(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to get token: %w", err)
		}
		//WithAuthToken移到getClient中
		var restClientWithAuth = restClient.WithAuthToken(token)
		return restClientWithAuth, nil // closing over client
	}

	getGQLClient := func(ctx context.Context) (*githubv4.Client, error) {
		token, err := cfg.Token(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to get token: %w", err)
		}

		var httpClient = &http.Client{
			//bearerAuthTransport移到getGQLClient处理
			Transport: &bearerAuthTransport{
				transport: http.DefaultTransport,
				token:     token,
			},
		}
		var gqlClientWithAuth = gqlClient.WithClient(httpClient)
		//bearerAuthTransport移到getGQLClient处理
		return githubv4.NewEnterpriseClient(gqlClientWithAuth.url, gqlClientWithAuth.httpClient), nil // closing over client
	}

	getRawClient := func(ctx context.Context) (*raw.Client, error) {
		client, err := getClient(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to get GitHub client: %w", err)
		}
		return raw.NewClient(client, apiHost.rawURL), nil // closing over client
	}

	// Create default toolsets
	tsg := github.DefaultToolsetGroup(cfg.ReadOnly, getClient, getGQLClient, getRawClient, cfg.Translator)
	err = tsg.EnableToolsets(enabledToolsets)

	if err != nil {
		return nil, fmt.Errorf("failed to enable toolsets: %w", err)
	}

	// Register all mcp functionality with the server
	tsg.RegisterAll(ghServer)

	if cfg.DynamicToolsets {
		dynamic := github.InitDynamicToolset(ghServer, tsg, cfg.Translator)
		dynamic.RegisterTools(ghServer)
	}

	return ghServer, nil
}
