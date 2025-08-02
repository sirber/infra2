package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Version information. These can be set at build time using -ldflags.
var (
	Version = "dev"
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

	// Read config
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

func showHelp() {
	fmt.Println("Usage: infra <action>")
	fmt.Println()
	fmt.Println("Available actions:")
	fmt.Println("  up      - Start the infrastructure")
	fmt.Println("  down    - Stop the infrastructure")
	fmt.Println("  pull    - Pull latest images")
	fmt.Println("  backup  - Backup data")
	fmt.Println("  restart - Restart the infrastructure")
	fmt.Println("  version - Show version information")
}

func printVersion() {
	fmt.Println(Version)
}

func main() {
	fmt.Println("Docker Self-Hosted Infra")
	fmt.Printf("Version: ")
	printVersion()
	fmt.Println()

	configPath := "infra.json"
	cfg, err := ensureConfig(configPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading config: %v\n", err)
		os.Exit(1)
	}

	if len(os.Args) < 2 {
		showHelp()
		return
	}

	action := os.Args[1]
	switch action {
	case "up":
		fmt.Println("Starting infrastructure...")
		dockerUp()

	case "down":
		fmt.Println("Stopping infrastructure...")
		dockerDown()

	case "pull":
		fmt.Println("Pulling latest images...")
		dockerPull()

	case "backup":
		fmt.Println("Backing up data...")
		dockerDown()
		backupData(cfg.ArchiveName)
		dockerUp()

	case "restart":
		fmt.Println("Restarting infrastructure...")
		dockerRestart()

	case "version":
		printVersion()

	default:
		fmt.Printf("Unknown action: %s\n\n", action)
		showHelp()
	}
}
