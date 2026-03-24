package cmd

import (
	"fmt"
	"yact/pkg/mihomo"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(configCmd)
}

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Hello from yact config",
	Long:  `There will be configs..`,
	RunE: func(cmd *cobra.Command, args []string) error {
		body := mihomo.Get("configs")
		fmt.Println(string(body))
		return nil
	},
}
