package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path"
)

var (
	homeUser   = os.Getenv("HOME")
	configPath = path.Join(homeUser, ".dnscli")
	configFile = path.Join(configPath, "config.json")
	config     = readConfig()
)

// A Config struct keeps the credentials information
// stored as a JSON document on `configFile`.
type Config struct {
	Token  string
	Domain string
	Mail   string
	ApiURL string
}

func createConfigPath() {
	// Create the config directory.
	err := os.Mkdir(configPath, 0755)
	if err != nil {
		log.Fatal("pathConfig: ", err)
	}

	// Create an empty config file.
	c := Config{}

	emptyConfig, err := json.Marshal(c)
	if err != nil {
		log.Fatal("pathConfig-Marshal: ", err)
	}

	ioutil.WriteFile(configFile, emptyConfig, 0644)
}
