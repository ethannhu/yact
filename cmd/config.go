package cmd

import (
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(configCmd)
}

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Hello from yact config",
	Long:  `There will be configs..`,
	Run: func(cmd *cobra.Command, args []string) {
		client := &http.Client{
			Timeout: 0,
		}
		req, _ := http.NewRequest("GET", "http://localhost:9097/configs", nil)
		req.Header.Add("Authorization", "Bearer 123456")

		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Req failed", err)
			return
		}
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		fmt.Println(string(body))
	},
}
