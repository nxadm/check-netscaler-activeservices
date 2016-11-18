package main

import "testing"

func TestRetrieveValues(t *testing.T) {
	file := "testfile_config.yaml"
	config, err := retrieveValues(file)
	if err != nil {
		t.Error("Can not parse configuration.")
	}
	if config.User != "some_user" {
		t.Error("Unexepected user value. Expected \"some_user\", got \"" + config.User + "\"")
	}
	if config.Pass != "some_password" {
		t.Error("Unexepected user value. Expected \"some_password\", got \"" + config.Pass + "\"")
	}
}
