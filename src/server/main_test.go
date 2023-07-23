package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

var jwt_token string

func init() {
	go func() { // Initialize the test server
		main()
	}()

	//generate token
	setupDotenv()
	var err error
	jwt_token, err = GenerateToken("testing")
	if err != nil {
		panic(err)
	}

}

func TestRegister(t *testing.T) {
	// Prepare the test request body
	requestBody := SystemInfo{
		Hostname:     "TestDevice",
		OS:           "Linux",
		Architecture: "x86_64",
	}
	requestBodyJSON, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatalf("Failed to marshal request body: %v", err)
	}

	// Create a new test request and response recorder
	req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewReader(requestBodyJSON))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+jwt_token)
	res, err := app.Test(req)
	if err != nil {
		t.Fatalf("Failed to perform test request: %v", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			t.Fatalf("Failed to close response body: %v", err)
		}
	}(res.Body)

	// Check the response status code
	if res.StatusCode != http.StatusOK {
		t.Fatalf("Expected status code %d but got %d", http.StatusOK, res.StatusCode)
	}

	// Parse the response JSON
	var responseBody map[string]int8
	err = json.NewDecoder(res.Body).Decode(&responseBody)
	if err != nil {
		t.Fatalf("Failed to parse response JSON: %v", err)
	}

	// Check if the response contains the expected "id" field
	if _, ok := responseBody["id"]; !ok {
		t.Fatalf("Response JSON does not contain \"id\" field")
	}

	var expected int8 = 3
	if responseBody["id"] != expected {
		t.Fatalf("Expected id %d but got %d", expected, responseBody["id"])
	}
}

func TestPing(t *testing.T) {
	// Create a new test request and response recorder
	req := httptest.NewRequest(http.MethodPost, "/ping/1", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+jwt_token)
	res, err := app.Test(req)
	if err != nil {
		t.Fatalf("Failed to perform test request: %v", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			t.Fatalf("Failed to close response body: %v", err)
		}
	}(res.Body)

	// Check the response status code
	if res.StatusCode != http.StatusOK {
		t.Fatalf("Expected status code %d but got %d", http.StatusOK, res.StatusCode)
	}

	// Parse the response JSON
	var responseBody map[string]string
	err = json.NewDecoder(res.Body).Decode(&responseBody)
	if err != nil {
		t.Fatalf("Failed to parse response JSON: %v", err)
	}

	// Check if the response contains the expected "message" field
	if _, ok := responseBody["message"]; !ok {
		t.Fatalf("Response JSON does not contain \"message\" field")
	}

	expected := "Pong!"
	if responseBody["message"] != expected {
		t.Fatalf("Expected message %s but got %s", expected, responseBody["message"])
	}
}
