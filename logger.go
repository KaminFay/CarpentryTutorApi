package main

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

// ErrorLevel is an enumerated data type to allow us to in code track the value of the errors we are passing in.
type ErrorLevel int

// These directly match the values for the logrus logging.
const (
	PANIC ErrorLevel = 1
	FATAL ErrorLevel = 2
	ERROR ErrorLevel = 3
	WARN  ErrorLevel = 4
	INFO  ErrorLevel = 5
	DEBUG ErrorLevel = 6
	TRACE ErrorLevel = 7
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func warningLog(logLocation string, logErr string) {
	log.WithFields(log.Fields{
		"Location": logLocation,
		"Error":    logErr,
	}).Warn("Something has gone wrong.")
}

func serverShutdownLog() {
	log.Warnln("Shutting down the server!")
}

func serverStartupLog() {
	log.Infoln("Starting up the server now!")
}

func testDatabaseConnectionLog(file string, logLocation string, logDescription string, logLevel ErrorLevel) {
	switch logLevel {
	case INFO:
		log.WithFields(log.Fields{
			"File":             file,
			"FunctionLocation": logLocation,
		}).Info(logDescription)
	case PANIC:
		log.WithFields(log.Fields{
			"File":             file,
			"FunctionLocation": logLocation,
		}).Info(logDescription)
	default:
		fmt.Println("Hitting default for some reason")
	}
}

func createUserLog() {

}

func getUserLog(file string, logLocation string, logDescription string, logLevel ErrorLevel) {
	switch logLevel {
	case INFO:
		log.WithFields(log.Fields{
			"File":             file,
			"FunctionLocation": logLocation,
		}).Info(logDescription)
	case PANIC:
		log.WithFields(log.Fields{
			"File":             file,
			"FunctionLocation": logLocation,
		}).Info(logDescription)
	default:
		fmt.Println("Hitting default for some reason")
	}
}
