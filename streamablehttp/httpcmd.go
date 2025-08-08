package streamablehttp

import (
	"context"
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func tokenFromContext(ctx context.Context) (string, error) {
	auth, ok := ctx.Value(authKey{}).(string)
	if !ok {
		return "", fmt.Errorf("missing auth from Authorization header")
	}
	return auth, nil
}
func HttpCmdfactory(version string) *cobra.Command {
	var httpCmd = &cobra.Command{
		Use:   "http",
		Short: "Start streamable http server",
		Long:  `Start a server that communicates via http streams using JSON-RPC messages.`,
		RunE: func(_ *cobra.Command, _ []string) error {

			// If you're wondering why we're not using viper.GetStringSlice("toolsets"),
			// it's because viper doesn't handle comma-separated values correctly for env
			// vars when using GetStringSlice.
			// https://github.com/spf13/viper/issues/380
			var enabledToolsets []string
			if err := viper.UnmarshalKey("toolsets", &enabledToolsets); err != nil {
				return fmt.Errorf("failed to unmarshal toolsets: %w", err)
			}

			httpServerConfig := HttpServerConfig{
				Address: viper.GetString("address"),
				Version: version,
				Host:    viper.GetString("host"),
				Token: func(ctx context.Context) (string, error) {

					auth, err := tokenFromContext(ctx)
					if err != nil {

						token := viper.GetString("personal_access_token")
						if token == "" {
							return "", errors.New("GITHUB_PERSONAL_ACCESS_TOKEN not set or missing auth from Authorization header")
						}
						return token, nil
					}
					return auth, nil
				},
				EnabledToolsets:      enabledToolsets,
				DynamicToolsets:      viper.GetBool("dynamic_toolsets"),
				ReadOnly:             viper.GetBool("read-only"),
				ExportTranslations:   viper.GetBool("export-translations"),
				EnableCommandLogging: viper.GetBool("enable-command-logging"),
				LogFilePath:          viper.GetString("log-file"),
			}
			return RunhttpServer(httpServerConfig)
		},
	}
	return httpCmd
}
