package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Record struct {
	Content      string
	DomainID     int64 `json:"domain_id"`
	ID           int64
	Name         string
	RecordType   string `json:"record_type"`
	SystemRecord bool   `json:"system_record"`
}

type MultipleRecords struct {
	Record Record
}

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
	r := setHeaders("GET", url)
	httpClient := http.Client{}

	response, err := httpClient.Do(r)
	if err != nil {
		log.Fatal("HTTPClient: ", err)
	}
	defer response.Body.Close()

	dataResponse := new([]MultipleRecords)
	err = json.NewDecoder(response.Body).Decode(&dataResponse)
	if err != nil {
		log.Fatal("Decode-Body: ", err)
	}

	//TODO: Filter and process response
	fmt.Println(dataResponse)
}
