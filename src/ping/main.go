package main

import (
	"fmt"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"time"
)

const (
	interval       = 10 * time.Minute
	maxRetries     = 3
	requestTimeout = 60 * time.Second
)

var url string

func main() {
	log.Info("Setting up dotenv")
	setupDotenv()
	url = os.Getenv("HEALTH_CHECK")
	if url == "" {
		log.Fatal("HEALTH_CHECK environment variable is not set")
	}

	setupHealthCheck()

	log.Info("Starting pinging: ", url)
	for {
		resp, err := pingTarget(url)
		if err != nil {
			log.Printf("Error pinging %s: %s\n", url, err)
		} else {
			log.Printf("Response from %s: %s\n", url, resp)
		}

		time.Sleep(interval)
	}
}

func setupHealthCheck() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		write, err := w.Write([]byte("ok"))
		if err != nil {
			log.Warn("Error writing response body: ", err)
		}
		if write != 2 {
			log.Warn("Error writing response body: ", err)
		}
	})
	go func() {
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			log.Fatal("Error starting health check server: ", err)
		}
	}()
}

func pingTarget(url string) (string, error) {
	client := &http.Client{
		Timeout: requestTimeout,
	}

	var resp *http.Response
	var err error
	for i := 0; i < maxRetries; i++ {
		resp, err = client.Get(url)
		if err == nil {
			break
		}
		log.Warnf("Error pinging %s: %s. Retrying...", url, err)
		time.Sleep(interval)
	}

	if err != nil {
		return "", fmt.Errorf("failed to ping %s after %d retries", url, maxRetries)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Warnf("Error closing response body: %s", err)
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected response status: %d", resp.StatusCode)
	}

	// Assuming the response body is textual, you can change this based on your API.
	// Read the response body and convert it to a string.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %s", err)
	}

	return string(body), nil
}

func setupDotenv() {
	err := godotenv.Load("src/.env")
	if err != nil {
		log.Warn("Error loading .env file: ", err)
		log.Warn("This is not a problem if you are running the server in production")
	}
}
