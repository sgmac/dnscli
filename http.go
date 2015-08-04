package main

import (
	"io"
	"log"
	"net/http"
)

func setHeaders(method, url string, data io.Reader) *http.Request {
	// Compose url and headers
	dnsHeader := config.Mail + ":" + config.Token

	r, err := http.NewRequest(method, url, data)
	if err != nil {
		log.Fatal("setHeaders-NewRequest ", err)
	}

	r.Header.Add("X-DNSimple-Token", dnsHeader)
	r.Header.Add("Accept", "application/json")
	r.Header.Add("Content-Type", "application/json")
	return r
}
