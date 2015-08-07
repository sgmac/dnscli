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

type ValidationRecord struct {
	Errors map[string][]string
}

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

// Update record provides two different JSON responses
// for both error and success. Capture the error in case
// the IP is invalid. It does not validate the IP itself.
func validateRecordUpdate(data []byte) {
	validation := new(ValidationRecord)

	err := json.Unmarshal(data, &validation)
	if err != nil {
		log.Fatal("json-unmarshal", err)
	}

	if len(validation.Errors["content"]) > 1 {
		fmt.Println(validation.Errors["content"][1])
	}
}

func stdoutHeader() {
	headerFmt := fmt.Sprintf("%s%21s%22s%27s%21s", "Type", "Name", "TTL", "RecordID", "Content")
	fmt.Println(headerFmt)
}

// Print out a record according the leftpadding
func stdoutRecord(r Record) {
	recordType := "%" + fmt.Sprintf("%d", (leftPadding-len(r.RecordType))+len(r.Name+config.Domain)) + "s"
	recordName := "%" + fmt.Sprintf("%d", (leftPadding-len(r.Name+config.Domain))+len(strconv.Itoa(r.TTL))) + "d"
	recordTTL := "%" + fmt.Sprintf("%d", (leftPadding-len(strconv.Itoa(r.TTL)))+len(strconv.Itoa(r.ID))) + "d"
	recordId := "%" + fmt.Sprintf("%d", (leftPadding-len(strconv.Itoa(r.ID)))+len(r.Content)) + "s"

	data := fmt.Sprint("%s", recordType, recordName, recordTTL, recordId)
	formatEnd := fmt.Sprintf("%s", data)
	fmt.Printf(fmt.Sprintf(formatEnd, r.RecordType, r.Name+"."+config.Domain, r.TTL, r.ID, r.Content))
	fmt.Println("")
}

func stdoutAutoRenew(a AutoRenew) {
	lo := "false"
	au := "false"
	if a.Lockable {
		lo = "true"
	}
	if a.AutoRenew {
		au = "true"
	}

	headerFmt := fmt.Sprintf("%s%24s%23s", "Domain", "Lockable", "AutoRenew")
	fmt.Println(headerFmt)

	dom := "%" + fmt.Sprintf("%d", (leftPadding-len(a.Domain))+len(lo)) + "s"
	lock := "%" + fmt.Sprintf("%d", (leftPadding-len(lo))+len(au)) + "s"
	data := fmt.Sprint("%s", dom, lock)
	formatEnd := fmt.Sprintf("%s", data)

	fmt.Printf(fmt.Sprintf(formatEnd, a.Domain, lo, au))
	fmt.Println("")
}
