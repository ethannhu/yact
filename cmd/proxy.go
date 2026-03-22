package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/spf13/cobra"
	"github.com/tidwall/gjson"
)

func init() {
	proxyQueryCmd.Flags().StringP("name", "n", "", "name of the proxy")
	if err := proxyQueryCmd.MarkFlagRequired("name"); err != nil {
		panic(err)
	}

	proxySetCmd.Flags().StringP("name", "n", "", "name of the proxy")

	if err := proxySetCmd.MarkFlagRequired("name"); err != nil {
		panic(err)
	}
	proxySetCmd.Flags().StringP("provider", "p", "", "name of the provider")
	if err := proxySetCmd.MarkFlagRequired("provider"); err != nil {
		panic(err)
	}

	proxyCmd.AddCommand(proxyListCmd, proxyQueryCmd, proxySetCmd)
	rootCmd.AddCommand(proxyCmd)
}

var proxyCmd = &cobra.Command{
	Use:   "proxy",
	Short: "Hello from yact proxy",
	Long:  `There will be proxies..`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

var proxyListCmd = &cobra.Command{
	Use:   "list",
	Short: "Hello from yact proxy",
	Long:  `There will be proxies..`,
	RunE: func(cmd *cobra.Command, args []string) error {
		client := &http.Client{
			Timeout: 0,
		}
		raw, err := cmd.Flags().GetBool("raw")
		if err != nil {
			return err
		}
		// proxy := url.PathEscape("🔰 选择节点")
		// selection := map[string]string{"name": "🇭🇰 香港Y02 | IEPL"}
		// jsonData, _ := json.Marshal(selection)
		// urlString := proxy
		// req, _ := http.NewRequest(http.MethodPut, urlString, bytes.NewBuffer(jsonData))
		urlString := "http://localhost:9097/proxies"
		req, _ := http.NewRequest(http.MethodGet, urlString, nil)
		req.Header.Add("Authorization", "Bearer 123456")
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Req failed", err)
			return err
		}
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		if !raw {
			result := gjson.Get(string(body), "proxies")
			result.ForEach(func(key, value gjson.Result) bool {
				proxyType := gjson.Get(value.Raw, "type").Str
				if proxyType == "Selector" {
					fmt.Println(key.Str + "->" + gjson.Get(value.Raw, "now").Str)
				}
				return true // keep iterating
			})
		} else {
			fmt.Println(string(body))
		}
		return nil
	},
}

var proxyQueryCmd = &cobra.Command{
	Use:   "query",
	Short: "Hello from yact proxy",
	Long:  `There will be proxies..`,
	RunE: func(cmd *cobra.Command, args []string) error {
		client := &http.Client{
			Timeout: 0,
		}
		raw, err := cmd.Flags().GetBool("raw")
		if err != nil {
			return err
		}
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}
		urlString := "http://localhost:9097/proxies/" + url.PathEscape(name)
		req, _ := http.NewRequest(http.MethodGet, urlString, nil)
		req.Header.Add("Authorization", "Bearer 123456")
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Req failed", err)
			return err
		}
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		if !raw {
			allProviders := gjson.Get(string(body), "all")
			allProviders.ForEach(func(key, value gjson.Result) bool {
				fmt.Println(value.Str)
				return true
			})
			provider := gjson.Get(string(body), "now")
			fmt.Println("Now: " + provider.Str)

		} else {
			fmt.Println(string(body))
		}
		return nil
	},
}

var proxySetCmd = &cobra.Command{
	Use:   "set",
	Short: "Hello from yact proxy",
	Long:  `There will be proxies..`,
	RunE: func(cmd *cobra.Command, args []string) error {
		client := &http.Client{
			Timeout: 0,
		}
		raw, err := cmd.Flags().GetBool("raw")
		if err != nil {
			return err
		}
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}
		provider, err := cmd.Flags().GetString("provider")
		if err != nil {
			return err
		}
		selection := map[string]string{"name": provider}
		jsonData, _ := json.Marshal(selection)
		urlString := "http://localhost:9097/proxies/" + url.PathEscape(name)
		req, _ := http.NewRequest(http.MethodPut, urlString, bytes.NewBuffer(jsonData))
		req.Header.Add("Authorization", "Bearer 123456")
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Req failed", err)
			return err
		}
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		if !raw {
			fmt.Println(resp.Status)
		} else {
			fmt.Println(string(body))
		}
		return nil
	},
}
