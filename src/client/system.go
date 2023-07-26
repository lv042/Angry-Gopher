package main

import (
	log "github.com/sirupsen/logrus"
	"os"
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
