package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var inspectCmd = &cobra.Command{
	Use:   "inspect [path]",
	Short: "Statically inspect project files (Dockerfile, docker-compose.yml)",
	Long: `Scans the given directory for Dockerfiles and docker-compose.yml files,
performing static analysis to identify common anti-patterns or issues.`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		targetPath := "."
		if len(args) > 0 {
			targetPath = args[0]
		}

		fmt.Printf("Inspecting project at: %s\n\n", targetPath)

		// Check Dockerfile
		dockerfilePath := filepath.Join(targetPath, "Dockerfile")
		if content, err := os.ReadFile(dockerfilePath); err == nil {
			fmt.Println("✅ Found Dockerfile")
			analyzeDockerfile(string(content))
		} else {
			fmt.Println("⚠️  No Dockerfile found in the directory.")
		}

		fmt.Println("---")

		// Check docker-compose.yml
		composePath := filepath.Join(targetPath, "docker-compose.yml")
		if content, err := os.ReadFile(composePath); err == nil {
			fmt.Println("✅ Found docker-compose.yml")
			analyzeCompose(string(content))
		} else {
			// Fallback to docker-compose.yaml
			composeYamlPath := filepath.Join(targetPath, "docker-compose.yaml")
			if content, err := os.ReadFile(composeYamlPath); err == nil {
				fmt.Println("✅ Found docker-compose.yaml")
				analyzeCompose(string(content))
			} else {
				fmt.Println("⚠️  No docker-compose.yml or docker-compose.yaml found in the directory.")
			}
		}
	},
}

func analyzeDockerfile(content string) {
	lines := strings.Split(content, "\n")
	hasLatest := false
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "FROM ") && strings.Contains(trimmed, ":latest") {
			hasLatest = true
		}
	}

	if hasLatest {
		fmt.Println("   ❌ Anti-pattern: Using ':latest' tag in FROM statement. Consider pinning to a specific version.")
	} else {
		fmt.Println("   ✅ Base images look well-defined.")
	}
}

func analyzeCompose(content string) {
	lines := strings.Split(content, "\n")
	hasLatest := false
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "image:") && strings.Contains(trimmed, ":latest") {
			hasLatest = true
		}
	}

	if hasLatest {
		fmt.Println("   ❌ Anti-pattern: Using ':latest' tag in Compose image. Consider pinning to a specific version.")
	} else {
		fmt.Println("   ✅ Compose images look well-defined.")
	}
}

func init() {
	rootCmd.AddCommand(inspectCmd)
}
