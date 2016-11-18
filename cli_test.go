package main

import (
	"os"
	"testing"
)

var defaultsTest = Defaults{
	Author:     author,
	Warning:    warning,
	Critical:   critical,
	Insecure:   insecure,
	Percentage: percentage,
	TimeOut:    timeOut,
	Version:    version,
}

func TestGetParams(t *testing.T) {

	// Define the cli combinations to test
	okCliTests := make(map[string][]string)
	okCliTests["okMinimalCli"] = []string{"cmd", "-f", "some_file", "-u", "some_url"}
	okCliTests["okMaximalCli"] =
		[]string{"cmd", "-f", "some_file", "-u", "some_url", "-w", "1", "-c", "1", "-t", "15", "-p", "-i"}
	for _, cli := range okCliTests {
		os.Args = cli
		getParams(defaultsTest)
	}
}