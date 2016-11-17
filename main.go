package main

import (
	"fmt"
	"os"
	"strconv"
)

const OK = 0
const WARNING = 1
const CRITICAL = 2
const UNKNOWN = 3

const author = "Claudio Ramirez <pub.claudio@gmail.com>"
const warning = 0
const critical = 0
const insecure = false
const percentage = false
const timeOut = 10
const version = "0.1.0"

func main() {
	// Command line interface
	defaults := Defaults{
		Author:     author,
		Warning:    warning,
		Critical:   critical,
		Insecure:   insecure,
		Percentage: percentage,
		TimeOut:    timeOut,
		Version:    version,
	}
	params := getParams(defaults)
	config, err := retrieveValues(params.ConfigFile)
	//fmt.Printf("[DEBUG] Params: %v\n", params)
	if err != nil {
		fmt.Printf("[UNKNOWN] Error reading the configuration file: %v\n", err)
		os.Exit(UNKNOWN)
	}

	// Query the server
	answer, err := getJson(config, params)
	//fmt.Printf("[DEBUG] Answer: %v\n", answer)
	if err != nil {
		fmt.Printf("[UNKNOWN] Can not decode the server's answer: %v\n", err)
		os.Exit(UNKNOWN)
	}
	activeServices, err := strconv.Atoi(answer.Lbvserver[0].Activeservices)
	if err != nil {
		fmt.Printf("[UNKNOWN] Invalid server answer: %v\n", err)
		os.Exit(UNKNOWN)
	}
	totalServices, err := strconv.Atoi(answer.Lbvserver[0].Totalservices)
	if err != nil {
		fmt.Printf("[UNKNOWN] Invalid server answer: %v\n", err)
		os.Exit(UNKNOWN)
	}
	if totalServices == 0 {
		fmt.Println("[UNKNOWN] Service doesn't seem to be configured (totalservices == 0)\n")
		os.Exit(UNKNOWN)
	}

	// Create the warnings
	var w, c int
	if params.Percentage {
		if params.Warning == 0 {
			params.Warning = 1
		}
		if params.Critical == 0 {
			params.Critical = 1
		}
		w = int(
			float32(totalServices) *
				(float32(params.Warning) / float32(100)))
		c = int(float32(totalServices) *
			(float32(params.Critical) / float32(100)))
	} else {
		w = params.Warning
		c = params.Critical
	}

	switch {
	case activeServices <= c:
		fmt.Printf("[CRITICAL] Threshold (%d), Active (%d), Total (%d)\n",
			c, activeServices, totalServices)
		os.Exit(CRITICAL)
	case activeServices <= w:
		fmt.Printf("[WARNING] Threshold (%d), Active (%d), Total (%d)\n",
			w, activeServices, totalServices)
		os.Exit(WARNING)
	default:
		fmt.Printf("[OK] Threshold (w:%d,c:%d), Active (%d), Total (%d)\n",
			w, c, activeServices, totalServices)
		os.Exit(OK)
	}

}
