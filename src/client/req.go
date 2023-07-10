package main

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
)

func register(serverURL string) error {
	// Marshal the struct to JSON
	payload, err := json.Marshal(sysInfo)
	if err != nil {
		return fmt.Errorf("Error marshalling JSON: %v", err)
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
		return fmt.Errorf("Error sending HTTP request: %v", err)
	}

	// Print the HTTP response status code and body
	log.WithFields(log.Fields{
		"status_code": resp.StatusCode(),
		"response":    string(resp.Body()),
	}).Info("HTTP Response")

	return nil
}
