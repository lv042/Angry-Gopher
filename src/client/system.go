package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"os/exec"
	"runtime"
)

func getSystemInfo() SystemInfo {
	// Initialize the sysInfo variable with system information
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatalf("Error getting hostname: %v\n", err)
	}

	osName := runtime.GOOS
	arch := runtime.GOARCH

	sysInfo := SystemInfo{
		Hostname:     hostname,
		OS:           osName,
		Architecture: arch,
	}

	log.WithFields(log.Fields{
		"hostname":     sysInfo.Hostname,
		"os":           sysInfo.OS,
		"architecture": sysInfo.Architecture,
	}).Info("System Information Initialized")

	return sysInfo
}

func addToAutoStartup() {

}

func getExecutablePath() (string, error) {
	ex, err := os.Executable()
	if err != nil {
		return "", err
	}
	return ex, nil
}

func generatePlistContent(executablePath string) string {
	return fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
    <!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
    <plist version="1.0">
    <dict>
        <key>Label</key>
        <string>com.your_company.your_program_name</string>
        <key>ProgramArguments</key>
        <array>
            <string>%s</string>
        </array>
        <key>RunAtLoad</key>
        <true/>
        <key>KeepAlive</key>
        <true/>
    </dict>
    </plist>`, executablePath)
}

func loadLaunchDaemon() error {
	cmd := exec.Command("sudo", "launchctl", "load", "plistPath")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
