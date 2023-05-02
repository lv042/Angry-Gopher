package main

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime"

	"github.com/valyala/fasthttp"
)

type SystemInfo struct {
	Hostname     string `json:"hostname"`
	OS           string `json:"os"`
	Architecture string `json:"architecture"`
}

func main() {
	// Define the remote server URL
	serverURL := "https://example.com/api/systeminfo"

	// Create a new struct with system information
	hostname, _ := os.Hostname()
	osName := runtime.GOOS
	arch := runtime.GOARCH
	sysInfo := SystemInfo{
		Hostname:     hostname,
		OS:           osName,
		Architecture: arch,
	}

	// Marshal the struct to JSON
	payload, err := json.Marshal(sysInfo)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	// Send the JSON payload via HTTP POST request
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.Header.SetContentType("application/json")
	req.SetRequestURI(serverURL)
	req.SetBody(payload)

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	if err := fasthttp.Do(req, resp); err != nil {
		fmt.Println("Error sending HTTP request:", err)
		return
	}

	// Print the HTTP response status code and body
	fmt.Println("Response Status:", resp.Status())
	fmt.Println("Response Body:", string(resp.Body()))
}
