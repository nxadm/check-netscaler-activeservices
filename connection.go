package main

import (
	"crypto/tls"
	"encoding/json"
	"net/http"
	"time"
)

type Answer struct {
	Lbvserver []ServicesCount
}

type ServicesCount struct {
	Activeservices string
	Totalservices  string
}

func getJson(config Config, params Params) (*Answer, error) {
	answer := &Answer{}
	client := &http.Client{Timeout: time.Second * time.Duration(params.TimeOut)}

	/* Skip certificate validation if required */
	if params.Insecure {
		client.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		}
	}

	/* Connect and convert JSON to datastructure */
	request, err := http.NewRequest("GET", params.URL, nil)
	request.SetBasicAuth(config.User, config.Pass)
	response, err := client.Do(request)
	if err != nil {
		return answer, err
	}
	defer response.Body.Close()
	err = json.NewDecoder(response.Body).Decode(answer)
	return answer, err
}
