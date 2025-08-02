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
		backupData("backup.tar.gz") // TODO: make archive name configurable
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
