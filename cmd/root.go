package cmd

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "yact",
	Short: "Hello from yact",
	Long:  `There will be proxies..`,
	Run: func(cmd *cobra.Command, args []string) {
		client := &http.Client{
			Timeout: 0,
		}
		req, _ := http.NewRequest("GET", "http://localhost:9097/logs?level=info", nil)
		req.Header.Add("Authorization", "Bearer 123456")

		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Req failed", err)
			return
		}
		defer resp.Body.Close()
		scanner := bufio.NewScanner(resp.Body)
		for scanner.Scan() {
			line := scanner.Text()
			fmt.Println("Received:", line)
		}
		if err := scanner.Err(); err != nil {
			log.Fatal("Read failed:", err)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
