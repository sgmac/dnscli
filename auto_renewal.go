package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type AutoRenew struct {
	ID        int    `json:"id"`
	Domain    string `json:"unicode_name"`
	Lockable  bool
	AutoRenew bool `json:"auto_renew"`
}

func enableAutoRenewal(domain string, enabled bool) {
	var url string
	var r *http.Request
	if domain != "" {
		config.Domain = domain
	} else if config.Domain == "" {
		fmt.Println("Set a domain in your configuration file or provide one.")
		os.Exit(1)
	}

	// Enable/Disable auto renewal
	if enabled {
		url = config.ApiURL + config.Domain + "/auto_renewal"
		r = setHeaders("POST", url, nil)
	} else {
		url = config.ApiURL + config.Domain + "/auto_renewal"
		r = setHeaders("DELETE", url, nil)
	}

	httpClient := http.Client{}
	response, err := httpClient.Do(r)
	if err != nil {
		log.Fatal("AutoRenewal-HTTPClient: ", err)
	}
	defer response.Body.Close()

	dataResponse := make(map[string]AutoRenew)
	err = json.NewDecoder(response.Body).Decode(&dataResponse)
	if err != nil {
		log.Fatal("enableAutoRenewal-Decode: ", err)
	}
	//fmt.Println(dataResponse)
	stdoutAutoRenew(dataResponse["domain"])
}
