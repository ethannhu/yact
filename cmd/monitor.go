package cmd

import (
	"bufio"
	"fmt"
	"log"
	"net/http"

	"charm.land/lipgloss/v2"
	"github.com/spf13/cobra"
)

func init() {
	trafficCommand.Flags().Bool("draw", false, "draw a panel")
	rootCmd.AddCommand(logCommand, trafficCommand, memoryCommand, connectionCommand)
}

var logCommand = &cobra.Command{
	Use:   "log",
	Short: "Hello from log",
	Long:  `There will be logs..`,
	RunE: func(cmd *cobra.Command, args []string) error {
		client := &http.Client{
			Timeout: 0,
		}
		req, _ := http.NewRequest("GET", "http://localhost:9097/logs?level=info", nil)
		req.Header.Add("Authorization", "Bearer 123456")

		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Req failed", err)
			return err
		}
		defer resp.Body.Close()
		scanner := bufio.NewScanner(resp.Body)
		for scanner.Scan() {
			line := scanner.Text()
			fmt.Println("Received:", line)
		}
		if err := scanner.Err(); err != nil {
			return err
		}
		return nil
	},
}

var trafficCommand = &cobra.Command{
	Use:   "traffic",
	Short: "Hello from traffic",
	Long:  `There will be traffic..`,
	RunE: func(cmd *cobra.Command, args []string) error {
		drawFlag, err := cmd.Flags().GetBool("draw")
		if err != nil {
			return err
		}
		client := &http.Client{
			Timeout: 0,
		}
		req, _ := http.NewRequest("GET", "http://localhost:9097/traffic", nil)
		req.Header.Add("Authorization", "Bearer 123456")

		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Req failed", err)
			return err
		}
		defer resp.Body.Close()
		scanner := bufio.NewScanner(resp.Body)

		if drawFlag == false {

			for scanner.Scan() {
				line := scanner.Text()
				fmt.Println("Received:", line)
			}
			if err := scanner.Err(); err != nil {
				log.Fatal("Read failed:", err)
			}
			return nil
		} else {
			var style = lipgloss.NewStyle().Bold(true).Background(lipgloss.Cyan)

			for scanner.Scan() {
				line := scanner.Text()
				lipgloss.Printf("\r%s\033[K", style.Render(line))
			}
			if err := scanner.Err(); err != nil {
				log.Fatal("Read failed:", err)
			}
			return nil
		}
	},
}

var memoryCommand = &cobra.Command{
	Use:   "memory",
	Short: "Hello from memory",
	Long:  `There will be traffic..`,
	Run: func(cmd *cobra.Command, args []string) {
		client := &http.Client{
			Timeout: 0,
		}
		req, _ := http.NewRequest("GET", "http://localhost:9097/memory", nil)
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

var connectionCommand = &cobra.Command{
	Use:   "connection",
	Short: "Hello from connection",
	Long:  `There will be traffic..`,
	Run: func(cmd *cobra.Command, args []string) {
		client := &http.Client{
			Timeout: 0,
		}
		req, _ := http.NewRequest("GET", "http://localhost:9097/connections", nil)
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
