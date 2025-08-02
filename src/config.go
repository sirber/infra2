package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Config structure for infra.json
type Config struct {
	ArchiveName string `json:"archive_name"`
}

func ensureConfig(path string) (Config, error) {
	var cfg Config
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return cfg, fmt.Errorf("%s not found", path)
	}
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return cfg, err
	}
	err = json.Unmarshal(data, &cfg)
	if err != nil {
		return cfg, err
	}
	return cfg, nil
}

func createDefaultConfig(path string) error {
	defaultCfg := Config{ArchiveName: "backup.tar.gz"}
	data, err := json.MarshalIndent(defaultCfg, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path, data, 0644)
}

func initConfig(configPath string) {
	if _, err := os.Stat(configPath); err == nil {
		fmt.Printf("%s already exists.\n", configPath)
		return
	}
	err := createDefaultConfig(configPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating config: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Created default config: %s\n", configPath)
}
