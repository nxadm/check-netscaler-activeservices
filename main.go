package main

import (
//"crypto/tls"
//"encoding/json"
//"fmt"
//"log"
//"net/http"
//"time"
)
import (
	"fmt"
	"os"
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
	if err != nil {
		fmt.Printf("Error reading the configuration file: %v\n", err)
		os.Exit(UNKNOWN)
	}
	fmt.Printf("%v\n", config)
	// Config file
	//answer := &Answer{}
	//err := getJson(url, answer)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("Answer: %v", answer)

}

//func getJson(url string, target interface{}) error {
//	client := &http.Client{Timeout: time.Second * timeOut}
//
//	// Skip certificate validation if required
//	if insecure {
//		client.Transport = &http.Transport{
//			TLSClientConfig: &tls.Config{
//				InsecureSkipVerify: true,
//			},
//		}
//	}
//
//	// Connect
//	request, err := http.NewRequest("GET", url, nil)
//	request.SetBasicAuth(user, pass)
//	response, err := client.Do(request)
//	if err != nil {
//		return err
//	}
//	defer response.Body.Close()
//
//	return json.NewDecoder(response.Body).Decode(target)
//}
