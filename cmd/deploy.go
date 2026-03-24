package cmd

import (
	"embed"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

//go:embed deploy.sh
var deployFS embed.FS

func init() {
	deployCmd.Flags().Bool("raw", false, "output script to stdout instead of writing to file")
	deployCmd.Flags().String("core", "mihomo.gz", "path to core archive")
	deployCmd.Flags().String("config", "config.yaml", "path to config file")
	rootCmd.AddCommand(deployCmd)
}

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Generate deploy script for mihomo installation",
	Long:  `Generate a deploy.sh script that can be used to install mihomo on target systems.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		raw, _ := cmd.Flags().GetBool("raw")
		corePath, _ := cmd.Flags().GetString("core")
		configPath, _ := cmd.Flags().GetString("config")

		// Expand ~ in config path if present
		configPath = expandHome(configPath)

		// Read embedded deploy.sh
		scriptBytes, err := deployFS.ReadFile("deploy.sh")
		if err != nil {
			return fmt.Errorf("failed to read embedded deploy.sh: %w", err)
		}

		// Replace template variables
		script := string(scriptBytes)
		script = strings.ReplaceAll(script, "{{.Core}}", corePath)
		script = strings.ReplaceAll(script, "{{.Config}}", configPath)

		if raw {
			fmt.Println(script)
		} else {
			err := os.WriteFile("deploy.sh", []byte(script), 0755)
			if err != nil {
				return fmt.Errorf("failed to write deploy.sh: %w", err)
			}
			fmt.Println("deploy.sh generated successfully")
		}
		return nil
	},
}

// expandHome expands ~ to the user's home directory
func expandHome(path string) string {
	if strings.HasPrefix(path, "~/") {
		home, err := os.UserHomeDir()
		if err == nil {
			return strings.Replace(path, "~", home, 1)
		}
	}
	return path
}
