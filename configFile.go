package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct { // Interface with the yaml configuration file
	User string `yaml:"user,omitempty"`
	Pass string `yaml:"pass,omitempty"`
}

func retrieveValues(file string) (Config, error) {
	c := Config{}

	// Read it
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return c, err
	}

	// Unmarshall it
	err = yaml.Unmarshal(data, &c)
	return c, err
}

func printSampleConfig() {
	sampleConf :=
`---
### check-netscaler-activeservices configuration ###
user: "ccis_readonly"
pass: "some_string"
`
	fmt.Println(sampleConf)
}
