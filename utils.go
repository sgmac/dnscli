package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
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
