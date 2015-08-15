package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var (
	urlRecords = config.ApiURL + "/domains/" + config.Domain + "/records/"
)

// Struct holding a single record response.
type Record struct {
	Content      string `json:"content"`
	DomainID     int    `json:"domain_id"`
	ID           int
	Name         string `json:"name"`
	RecordType   string `json:"record_type"`
	SystemRecord bool   `json:"system_record"`
	TTL          int
}

// Listing records returns a slice of records.
type MultipleRecords []map[string]Record

func listRecordsDomain(domain string) {
	// If domain is empty, use the value from the cli.
	isDomainEmpty(domain)

	r := setHeaders("GET", urlRecords, nil)
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
	isDomainEmpty(domain)

	updateContent := make(map[string]string)
	updateContent["content"] = content
	data, err := json.Marshal(updateContent)
	d := strings.NewReader(string(data))

	url := urlRecords + id
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

func getRecordDomain(domain, id string) {
	isDomainEmpty(domain)

	url := urlRecords + id
	r := setHeaders("GET", url, nil)
	httpClient := http.Client{}

	response, err := httpClient.Do(r)
	if err != nil {
		log.Fatal("HTTPClient: ", err)
	}
	defer response.Body.Close()

	dataResponse := make(map[string]Record)
	err = json.NewDecoder(response.Body).Decode(&dataResponse)
	if err != nil {
		log.Fatal("getRecordDomain-Decode: ", err)
	}
	stdoutHeader()
	stdoutRecord(dataResponse["record"])
}

func deleteRecordDomain(domain, id string) {
	isDomainEmpty(domain)

	url := urlRecords + id
	r := setHeaders("DELETE", url, nil)
	httpClient := http.Client{}

	response, err := httpClient.Do(r)
	if err != nil {
		log.Fatal("HTTPClient: ", err)
	}
	defer response.Body.Close()

	dataResponse := make(map[string]string)
	err = json.NewDecoder(response.Body).Decode(&dataResponse)
	if err != nil {
		log.Fatal("deleteRecordDomain-Decode: ", err)
	}

	// In case the record ID does not exist
	if msg, ok := dataResponse["message"]; ok {
		fmt.Println("", msg)
	}
}

func createRecordDomain(domain, name, content, recordType string) {
	isDomainEmpty(domain)

	re := make(map[string]Record)
	record := Record{}
	record.Content = content
	record.RecordType = recordType
	record.Name = name
	re["record"] = record

	dataJSON, err := json.Marshal(re)
	data := strings.NewReader(string(dataJSON))
	if err != nil {
		log.Fatal("createRecordDomain-Marshal: ", err)
	}

	r := setHeaders("POST", urlRecords, data)
	httpClient := http.Client{}

	response, err := httpClient.Do(r)
	if err != nil {
		log.Fatal("HTTPClient: ", err)
	}
	defer response.Body.Close()

	dataResponse := make(map[string]Record)
	resp, err := ioutil.ReadAll(response.Body)

	// JSON response differs for success, failure
	// and record already defined.
	validateRecordUpdate(resp)
	isRecordDefined(resp)
	err = json.NewDecoder(strings.NewReader(string(resp))).Decode(&dataResponse)
	if err != nil {
		log.Fatal("createRecordDomain-Decode: ", err)
	}

	stdoutHeader()
	stdoutRecord(dataResponse["record"])
}
