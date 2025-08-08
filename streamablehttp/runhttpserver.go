package streamablehttp

import (
	"context"
	"fmt"
	"gitee.com/masx200/github-mcp-server/pkg/errors"
	mcplog "gitee.com/masx200/github-mcp-server/pkg/log"
	"gitee.com/masx200/github-mcp-server/pkg/translations"
	"github.com/mark3labs/mcp-go/server"
	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

func RunhttpServer(cfg HttpServerConfig) error {
	// Create app context
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	t, dumpTranslations := translations.TranslationHelper()

	logrusLogger := logrus.New()
	if cfg.Pretty {

		logrusLogger.SetFormatter(&logrus.JSONFormatter{PrettyPrint: true})
	} else {
		logrusLogger.SetFormatter(&prefixed.TextFormatter{
			ForceFormatting: true,
			ForceColors:     true,
			FullTimestamp:   true,
			// 想继续输出 JSON 可改用 prettyjson-formatter 等
		})
	}
	logrusLogger.SetOutput(os.Stderr)
	if cfg.LogFilePath != "" {
		file, err := os.OpenFile(cfg.LogFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
		if err != nil {
			return fmt.Errorf("failed to open log file: %w", err)
		}

		logrusLogger.SetLevel(logrus.DebugLevel)
		logrusLogger.SetOutput(file)
	}
	stdLogger := log.New(logrusLogger.Writer(), "httpserver", 0)
	var ghServer *server.MCPServer
	var err error

	ghServer, err = NewMCPServer(MCPServerConfig{
		Version:         cfg.Version,
		Host:            cfg.Host,
		Token:           cfg.Token,
		EnabledToolsets: cfg.EnabledToolsets,
		DynamicToolsets: cfg.DynamicToolsets,
		ReadOnly:        cfg.ReadOnly,
		Translator:      t,
	}, CreateHooksWithEventLogging(logrusLogger, cfg.Pretty))

	// } else {
	// 	ghServer, err = NewMCPServer(MCPServerConfig{
	// 		Version:         cfg.Version,
	// 		Host:            cfg.Host,
	// 		Token:           cfg.Token,
	// 		EnabledToolsets: cfg.EnabledToolsets,
	// 		DynamicToolsets: cfg.DynamicToolsets,
	// 		ReadOnly:        cfg.ReadOnly,
	// 		Translator:      t,
	// 	})

	// }
	if err != nil {
		return fmt.Errorf("failed to create MCP server: %w", err)
	}

	httpServer := server.NewStreamableHTTPServer(ghServer,
		server.WithHTTPContextFunc(authFromRequest),
		server.WithLogger(NewLoggerAdapter(stdLogger)),
		server.WithHTTPContextFunc(func(ctx context.Context, r *http.Request) context.Context {

			return errors.ContextWithGitHubErrors(ctx)

		}),
	)

	// httpServer.SetLogger(stdLogger)

	if cfg.ExportTranslations {
		// Once server is initialized, all translations are loaded
		dumpTranslations()
	}

	// Start listening for messages
	errC := make(chan error, 1)
	go func() {
		in, out := io.Reader(os.Stdin), io.Writer(os.Stdout)

		if cfg.EnableCommandLogging {
			loggedIO := mcplog.NewIOLogger(in, out, logrusLogger)
			in, out = loggedIO, loggedIO
		}
		// enable GitHub errors in the context

		// errC <- httpServer.Listen(ctx, in, out)
		errC <- httpServer.Start(cfg.Address)
	}()

	var addr = cfg.Address
	fullAddr, err := NormalizeAddress(addr)
	if err != nil {
		fmt.Printf("Error parsing %s: %v\n", addr, err)
		return err
	}
	fmt.Printf("Normalized: %s\n", fullAddr)

	// Output github-mcp-server string
	_, _ = fmt.Fprintf(os.Stderr, "GitHub MCP Server running on http://%s\n", fullAddr)
	log.Println("GitHub MCP Server running on http://" + fullAddr)

	PrintGitHubEnvVars()

	// Wait for shutdown signal
	select {
	case <-ctx.Done():
		logrusLogger.Infof("shutting down server...")
	case err := <-errC:
		if err != nil {
			return fmt.Errorf("error running server: %w", err)
		}
	}

	return nil
}

// NormalizeAddress 将输入地址解析为标准化格式 "ip:port"
func NormalizeAddress(addr string) (string, error) {
	// 分割主机和端口
	host, port, err := net.SplitHostPort(addr)
	if err != nil {
		return "", fmt.Errorf("invalid address format: %w", err)
	}

	// 处理端口合法性
	if _, err := strconv.ParseUint(port, 10, 16); err != nil {
		return "", fmt.Errorf("invalid port: %s", port)
	}

	// 处理主机部分（IP或空）
	var ipStr string
	switch {
	case host == "":
		ipStr = "0.0.0.0" // 空IP表示监听所有接口
	default:
		// 尝试解析为IP地址
		ip := net.ParseIP(host)
		if ip == nil {
			return "", fmt.Errorf("invalid IP: %s", host)
		}
		ipStr = ip.String() // 标准化IP格式（如IPv6压缩表示）
	}

	// 组合为标准格式
	return net.JoinHostPort(ipStr, port), nil
}
