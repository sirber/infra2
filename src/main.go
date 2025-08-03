package main

import (
	"fmt"
	"os"
)

// Version information. These can be set at build time using -ldflags.
var (
	Version = "dev"
)

func showHelp() {
	fmt.Println("Usage: infra <action>")
	fmt.Println()
	fmt.Println("Available actions:")
	fmt.Println("  init    - Create default config file")
	fmt.Println("  up      - Start the infrastructure")
	fmt.Println("  start   - Stop running containers")
	fmt.Println("  stop    - Start running containers")
	fmt.Println("  down    - Stop the infrastructure")
	fmt.Println("  pull    - Pull latest images")
	fmt.Println("  backup  - Backup data")
	fmt.Println("  restart - Restart the infrastructure")
	fmt.Println("  status  - Show status of all enabled services")
	fmt.Println("  logs    - Show logs of all enabled services")
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
	case "start":
		fmt.Println("Starting infrastructure...")
		dockerStart()
	case "down":
		fmt.Println("Stopping infrastructure...")
		dockerDown()
	case "stop":
		fmt.Println("Stopping infrastructure...")
		dockerStop()
	case "pull":
		fmt.Println("Pulling latest images...")
		dockerPull()
	case "backup":
		dockerStop()
		fmt.Println("Backing up data...")
		backupData(cfg.ArchiveName)
		dockerStart()
	case "restart":
		fmt.Println("Restarting infrastructure...")
		dockerRestart()
	case "status":
		fmt.Println("Showing status of all enabled services...")
		dockerStatus()
	case "logs":
		fmt.Println("Showing logs of all enabled services...")
		dockerLogs()
	case "version":
		printVersion()
	case "init":
		initConfig(configPath)
		return
	default:
		fmt.Printf("Unknown action: %s\n\n", action)
		showHelp()
	}
}
