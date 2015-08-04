package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type Record struct {
	Content      string
	DomainID     int `json:"domain_id"`
	ID           int
	Name         string
	RecordType   string `json:"record_type"`
	SystemRecord bool   `json:"system_record"`
	TTL          int
}

type MultipleRecords []map[string]Record

func listRecordsDomain(domain string) {
	// Default domain is empty, use the value from the cli
	if domain != "" {
		config.Domain = domain
	} else if config.Domain == "" {
		fmt.Println("Set a domain in your configuration file or provide one.")
		os.Exit(1)
	}

	// Compose url and headers
	url := config.ApiURL + config.Domain + "/records/"
	r := setHeaders("GET", url, nil)
	httpClient := http.Client{}

	response, err := httpClient.Do(r)
	if err != nil {
		log.Fatal("HTTPClient: ", err)
	}
	defer response.Body.Close()

	dataResponse := MultipleRecords{}
	err = json.NewDecoder(response.Body).Decode(&dataResponse)
	if err != nil {
		log.Fatal("Decode-Body: ", err)
	}

	filterRecords(dataResponse)
}

func updateRecordDomain(domain, content, id string) {
	if domain != "" {
		config.Domain = domain
	} else if config.Domain == "" {
		fmt.Println("Set a domain in your configuration file or provide one.")
		os.Exit(1)
	}

	updateContent := make(map[string]string)
	updateContent["content"] = content
	data, err := json.Marshal(updateContent)
	d := strings.NewReader(string(data))

	url := config.ApiURL + config.Domain + "/records/" + id
	req := setHeaders("PUT", url, d)
	httpClient := http.Client{}

	response, err := httpClient.Do(req)
	if err != nil {
		log.Fatal("updateRecord-Do: ", err)
	}
	defer response.Body.Close()

	resp, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal("updateRecord-ReadAll: ", err)
	}
	validateRecordUpdate(resp)
}

func getRecordDomain(domain, record string) {
	if domain != "" {
		config.Domain = domain
	} else if config.Domain == "" {
		fmt.Println("Set a domain in your configuration file or provide one.")
		os.Exit(1)
	}
	// TODO: Probably best to move URLs out, vars or in the config file.
	url := config.ApiURL + config.Domain + "/records/" + record
	fmt.Println("url:", url)
	r := setHeaders("GET", url, nil)
	httpClient := http.Client{}

	response, err := httpClient.Do(r)
	if err != nil {
		log.Fatal("HTTPClient: ", err)
	}
	defer response.Body.Close()

	dataResponse := map[string]Record{}
	err = json.NewDecoder(response.Body).Decode(&dataResponse)
	if err != nil {
		log.Fatal("getRecordDomain-Decode: ", err)
	}
	stdoutHeader()
	stdoutRecord(dataResponse["record"])
}
