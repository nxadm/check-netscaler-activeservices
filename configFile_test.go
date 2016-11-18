package main

import "testing"

func TestRetrieveValues(t *testing.T) {
	file := "testfile_config.yaml"
	config, err := retrieveValues(file)
	if err != nil {
		t.Error("Can not parse configuration.")
	}
	if config.User != "someuser" {
		t.Error("Unexepected user value. Expected \"someuser\", got \"" + config.User + "\"")
	}
	if config.Pass != "somepassword" {
		t.Error("Unexepected user value. Expected \"somepassword\", got \"" + config.Pass + "\"")
	}
}
