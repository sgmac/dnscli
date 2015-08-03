package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
)

var (
	leftPadding int = 22
)

func readConfig() *Config {
	config := &Config{}
	configData, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatal("readConfig-ReadFile", err)
	}

	err = json.Unmarshal(configData, config)
	if err != nil {
		log.Fatal("readConfig-Unmarshal", err)
	}
	return config
}

func filterRecords(m MultipleRecords) {
	stdoutHeader()
	for _, item := range m {
		r := item["record"]
		// Skip system records are managed by DNSimple
		if !r.SystemRecord {
			stdoutRecord(r)
		}
	}
}

func stdoutHeader() {
	headerFmt := fmt.Sprintf("%s%21s%22s%26s", "Type", "Name", "TTL", "Content")
	fmt.Println(headerFmt)
}

// Print out a record according the leftpadding
func stdoutRecord(r Record) {
	recordType := "%" + fmt.Sprintf("%d", (leftPadding-len(r.RecordType))+len(r.Name+config.Domain)) + "s"
	recordName := "%" + fmt.Sprintf("%d", (leftPadding-len(r.Name+config.Domain))+len(strconv.Itoa(r.TTL))) + "d"
	recordTTL := "%" + fmt.Sprintf("%d", (leftPadding-len(strconv.Itoa(r.TTL)))+len(r.Content)) + "s"
	data := fmt.Sprint("%s", recordType, recordName, recordTTL)
	formatEnd := fmt.Sprintf("%s", data)
	fmt.Printf(fmt.Sprintf(formatEnd, r.RecordType, r.Name+"."+config.Domain, r.TTL, r.Content))
	fmt.Println("")
}
