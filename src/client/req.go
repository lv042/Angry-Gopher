// Function to generate a JWT token using the provided API and return the ID
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"strconv"
)

// Function to generate a JWT token using the provided API and return the ID or an error
func register(baseURL string, token string, sysInfo *SystemInfo) (int, error) {
	apiEndpoint := baseURL + "/register"

	payload, err := json.Marshal(sysInfo)
	if err != nil {
		return 0, err
	}

	// Prepare the request
	req, err := http.NewRequest("POST", apiEndpoint, bytes.NewBuffer(payload))
	if err != nil {
		return 0, err
	}

	// Set the JWT token in the request header
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	// Make the request to the API
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(resp.Body)

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("API request failed with status code %d", resp.StatusCode)
	}

	// Parse the response and extract the ID
	var responseData map[string]int
	err = json.NewDecoder(resp.Body).Decode(&responseData)
	if err != nil {
		return 0, err
	}

	id, ok := responseData["id"]
	if !ok {
		return 0, fmt.Errorf("failed to get ID from API response")
	}

	// Log the success
	log.WithFields(log.Fields{
		"token": token,
		"id":    id,
	}).Info("Token generated and ID obtained from API")

	return id, nil
}

func getCommands(token string, id int) ([]CommandResult, error) {
	apiEndpoint := device.ServerURL + "/cmd/" + strconv.Itoa(id)

	// Prepare the request
	req, err := http.NewRequest("GET", apiEndpoint, nil)
	if err != nil {
		return nil, err
	}

	// Set the JWT token in the request header
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	// Make the request to the API
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Warn(err)
		}
	}(resp.Body)

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status code %d", resp.StatusCode)
	}

	// Parse the response and extract the ID
	var commandResults []CommandResult
	err = json.NewDecoder(resp.Body).Decode(&commandResults)
	if err != nil {
		return nil, err
	}

	return commandResults, nil
}

func postCommandResult(token string, id int, commandResult CommandResult) error {
	apiEndpoint := device.ServerURL + "/cmd/" + strconv.Itoa(id)

	// Prepare the request body
	requestBody, err := json.Marshal(commandResult)
	if err != nil {
		return err
	}

	// Prepare the request
	req, err := http.NewRequest("POST", apiEndpoint, bytes.NewReader(requestBody))
	if err != nil {
		return err
	}

	// Set the JWT token in the request header
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	// Make the request to the API
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Warn(err)
		}
	}(resp.Body)

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("API request failed with status code %d", resp.StatusCode)
	}

	return nil
}
