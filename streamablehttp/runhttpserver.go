package streamablehttp

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"gitee.com/masx200/github-mcp-server/pkg/errors"
	mcplog "gitee.com/masx200/github-mcp-server/pkg/log"
	"gitee.com/masx200/github-mcp-server/pkg/translations"
	"github.com/mark3labs/mcp-go/server"
	"github.com/sirupsen/logrus"
)

func RunhttpServer(cfg HttpServerConfig) error {
	// Create app context
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	t, dumpTranslations := translations.TranslationHelper()

	ghServer, err := NewMCPServer(MCPServerConfig{
		Version:         cfg.Version,
		Host:            cfg.Host,
		Token:           cfg.Token,
		EnabledToolsets: cfg.EnabledToolsets,
		DynamicToolsets: cfg.DynamicToolsets,
		ReadOnly:        cfg.ReadOnly,
		Translator:      t,
	})
	if err != nil {
		return fmt.Errorf("failed to create MCP server: %w", err)
	}

	logrusLogger := logrus.New()
	if cfg.LogFilePath != "" {
		file, err := os.OpenFile(cfg.LogFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
		if err != nil {
			return fmt.Errorf("failed to open log file: %w", err)
		}

		logrusLogger.SetLevel(logrus.DebugLevel)
		logrusLogger.SetOutput(file)
	}
	stdLogger := log.New(logrusLogger.Writer(), "httpserver", 0)
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

	// Output github-mcp-server string
	_, _ = fmt.Fprintf(os.Stderr, "GitHub MCP Server running on http://%s\n", cfg.Address)
	log.Println("GitHub MCP Server running on http://" + cfg.Address)
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
