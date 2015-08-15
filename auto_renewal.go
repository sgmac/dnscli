package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type AutoRenew struct {
	ID        int    `json:"id"`
	Domain    string `json:"unicode_name"`
	Lockable  bool
	AutoRenew bool `json:"auto_renew"`
}

func enableAutoRenewal(domain string, enabled bool) {
	isDomainEmpty(domain)

	var url string
	var r *http.Request

	// Enable/Disable auto renewal
	if enabled {
		url = config.ApiURL + "/domains/" + config.Domain + "/auto_renewal"
		r = setHeaders("POST", url, nil)
	} else {
		url = config.ApiURL + "/domains/" + config.Domain + "/auto_renewal"
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
	stdoutAutoRenew(dataResponse["domain"])
}
