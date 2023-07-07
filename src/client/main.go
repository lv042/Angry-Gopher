package main

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

// Define the server URL
var serverURL = "https://example.com/api/systeminfo"
var sysInfo *SystemInfo

type SystemInfo struct {
	Hostname     string `json:"hostname"`
	OS           string `json:"os"`
	Architecture string `json:"architecture"`
}

type CommandResult struct {
	Message string    `json:"message"`
	Time    time.Time `json:"time"`
	Dir     string    `json:"dir"`
}

func initSysInfo() {
	// Initialize the sysInfo variable with system information
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatalf("Error getting hostname: %v\n", err)
	}

	osName := runtime.GOOS
	arch := runtime.GOARCH

	sysInfo = &SystemInfo{
		Hostname:     hostname,
		OS:           osName,
		Architecture: arch,
	}

	log.WithFields(log.Fields{
		"hostname":     sysInfo.Hostname,
		"os":           sysInfo.OS,
		"architecture": sysInfo.Architecture,
	}).Info("System Information Initialized")
}

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

func main() {
	runCmd("ls -a")

	//time.Sleep(100 * time.Second)
	//// Initialize the logger
	//log.SetFormatter(&log.JSONFormatter{})
	//log.SetOutput(os.Stdout)
	//
	//// Initialize the sysInfo variable with system information
	//initSysInfo()
	//
	//// Register the system information with the server
	//if err := register(serverURL); err != nil {
	//	log.WithError(err).Error("Error registering system information")
	//	return
	//}
	//
	//log.Info("System Information Registered Successfully")

}

func runCmd(inputCmd string) CommandResult {
	//split up the command and arguments
	modInput := strings.Split(inputCmd, " ")

	cmd := exec.Command(modInput[0], modInput[1:]...)
	res, _ := cmd.CombinedOutput()
	dir, _ := exec.Command("pwd").CombinedOutput()

	//fmt.Println("Command: ", cmd.ProcessState)

	if cmd.ProcessState == nil {
		res = []byte("Could not run command")
	}

	commandResult := CommandResult{
		Message: string(res),
		Time:    time.Now(),
		Dir:     string(dir),
	}

	fmt.Println("Message: ", commandResult.Message, "\nTime: ", commandResult.Time, "\nDir: ", commandResult.Dir)

	return commandResult
}
