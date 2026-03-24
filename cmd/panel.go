package cmd

import (
	"bufio"
	"fmt"
	"log"
	"yact/pkg/mihomo"

	"charm.land/lipgloss/v2"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(logCommand, trafficCommand, memoryCommand, connectionCommand)
}

var styleInfo = lipgloss.NewStyle().Foreground(lipgloss.Green)
var styleWarn = lipgloss.NewStyle().Foreground(lipgloss.Red)

var logCommand = &cobra.Command{
	Use:   "log",
	Short: "Hello from log",
	Long:  `There will be logs..`,
	RunE: func(cmd *cobra.Command, args []string) error {
		resp := mihomo.GetStream("logs")
		scanner := bufio.NewScanner(resp.Body)
		defer resp.Body.Close()
		raw, _ := cmd.Flags().GetBool("raw")
		for scanner.Scan() {
			line := scanner.Text()
			if raw {
				fmt.Println(line)
			} else {
				msgType, message := mihomo.ParseLog(line)
				_ = msgType
				fmt.Printf("%s %s %s %s\n", message.Protocol, message.Source, message.Destination, message.Chain)
			}

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
		raw, _ := cmd.Flags().GetBool("raw")
		resp := mihomo.GetStream("traffic")
		scanner := bufio.NewScanner(resp.Body)
		defer resp.Body.Close()
		for scanner.Scan() {
			line := scanner.Text()
			if raw {
				fmt.Println(line)
			} else {
				up, down := mihomo.ParseTraffic(line)
				lipgloss.Printf("%d %d\n", up, down)
			}
		}
		if err := scanner.Err(); err != nil {
			log.Fatal("Read failed:", err)
		}
		return nil
	},
}

var memoryCommand = &cobra.Command{
	Use:   "memory",
	Short: "short memory",
	Long:  `long memory`,
	Run: func(cmd *cobra.Command, args []string) {
		raw, _ := cmd.Flags().GetBool("raw")
		resp := mihomo.GetStream("memory")
		defer resp.Body.Close()
		scanner := bufio.NewScanner(resp.Body)
		for scanner.Scan() {

			line := scanner.Text()
			if raw {
				fmt.Println(line)
			} else {
				inuse, oslimit := mihomo.ParseMemory(line)
				fmt.Printf("%d %d\n", inuse, oslimit)
			}

		}
		if err := scanner.Err(); err != nil {
			log.Fatal("Read failed:", err)
		}
	},
}

var connectionCommand = &cobra.Command{
	Use:   "connection",
	Short: "short connection",
	Long:  `long connection`,
	Run: func(cmd *cobra.Command, args []string) {
		raw, _ := cmd.Flags().GetBool("raw")
		line := string(mihomo.Get("connections"))
		if raw {
			fmt.Println(line)
		} else {
			connections := mihomo.ParseConnections(line)
			for _, c := range connections {
				fmt.Println(c.Host)
			}
		}
	},
}
