package main

import (
	"log"
	"os"
	"path"
	"testing"
)

func createFakeConfigPath() {
	// Change configPath
	err := os.Setenv("TMPDIR", "/tmp")
	if err != nil {
		log.Fatal(" ", err)
	}
	tmpPath := os.Getenv("TMPDIR")

	// Set a configuration path environment in /tmp/.dnscli
	configPath = path.Join(tmpPath, ".dnscli")
	configFile = path.Join(configPath, "config.json")
	createConfigPath()
}

func TestCreateConfigPath(t *testing.T) {
	createFakeConfigPath()
	// Clean up
	defer func() {
		os.RemoveAll(configPath)
	}()

	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		t.Error("configuration path not found")
	}
}
