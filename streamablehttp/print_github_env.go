package streamablehttp

import (
	"fmt"
	"os"
	"strings"
)

// PrintGitHubEnvVars prints all environment variables that start with "github" (case-insensitive)
func PrintGitHubEnvVars() {
	for _, env := range os.Environ() {
		parts := strings.SplitN(env, "=", 2)
		if len(parts) < 1 {
			continue
		}

		key := parts[0]
		if strings.HasPrefix(strings.ToLower(key), "github") {

			fmt.Fprintln(os.Stderr, env)

			//	fmt.Println(env)
		}
	}
}
