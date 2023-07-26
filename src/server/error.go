package main

import log "github.com/sirupsen/logrus"

func checkFatalError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}
