package mihomo

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

type Core struct {
	Address string
	Secret  string
}

func Get(path string) []byte {

	c := Core{
		Address: "http://localhost:9097/",
		Secret:  "123456",
	}
	client := &http.Client{
		Timeout: 0,
	}
	url := c.Address + path
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Add("Authorization", "Bearer "+c.Secret)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Req failed", err)
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	return body
}

func GetStream(path string) *http.Response {
	c := Core{
		Address: "http://localhost:9097/",
		Secret:  "123456",
	}
	client := &http.Client{
		Timeout: 0,
	}
	url := c.Address + path
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Add("Authorization", "Bearer "+c.Secret)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Req failed", err)
		panic(err)
	}
	return resp
}

func Put(path string, data []byte) string {

	c := Core{
		Address: "http://localhost:9097/",
		Secret:  "123456",
	}
	client := &http.Client{
		Timeout: 0,
	}
	url := c.Address + path
	req, _ := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(data))
	req.Header.Add("Authorization", "Bearer "+c.Secret)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Req failed", err)
		panic(err)
	}
	defer resp.Body.Close()
	return resp.Status
}
