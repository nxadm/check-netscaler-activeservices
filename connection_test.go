package main

import (
	"testing"
)

func TestGetJson(t *testing.T) {
	c := Config{User: "me", Pass: "secret"}
	p := Params{
		ConfigFile: "string",
		Critical:   0,
		Insecure:   true,
		Percentage: true,
		TimeOut:    1,
		URL:        "->string-<",
		Warning:    0,
	}

	answer, err := getJson(c, p)
	if err == nil {
		t.Error("This connection attempt must return an error.")
	}
	if answer.Lbvserver != nil {
		t.Error("This connection answer must be empty.")
	}
}
