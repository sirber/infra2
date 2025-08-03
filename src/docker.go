package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func findEnabledDirs(root string, all bool) ([]string, error) {
	var dirs []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			// Skip directories/files we can't access
			if os.IsPermission(err) {
				return filepath.SkipDir
			}
			return nil
		}

		if info.IsDir() == false {
			return nil
		}

		// Look for docker-compose.yml in the directory
		dockerPath := filepath.Join(path, "docker-compose.yml")
		if _, err := os.Stat(dockerPath); err != nil {
			return nil
		}

		// If 'all' is true, include the service regardless of 'enabled'
		if all {
			dirs = append(dirs, path)
		} else {
			enabledPath := filepath.Join(path, "enabled")
			if _, err := os.Stat(enabledPath); err == nil {
				dirs = append(dirs, path)
			}
		}

		return nil
	})
	return dirs, err
}

func runDockerComposeCmdInDirs(dirs []string, args ...string) {
	for _, dir := range dirs {
		cmd := exec.Command("docker", append([]string{"compose"}, args...)...)
		cmd.Dir = dir
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		fmt.Printf("Running 'docker compose %v' in %s\n", args, dir)
		if err := cmd.Run(); err != nil {
			fmt.Printf("Error running docker compose in %s: %v\n", dir, err)
		}
	}
}

func dockerUp() {
	dirs, err := findEnabledDirs(".", false)
	if err != nil {
		fmt.Printf("Error scanning directories: %v\n", err)
		os.Exit(1)
	}
	runDockerComposeCmdInDirs(dirs, "up", "-d")
}

func dockerDown() {
	dirs, err := findEnabledDirs(".", true)
	if err != nil {
		fmt.Printf("Error scanning directories: %v\n", err)
		os.Exit(1)
	}
	runDockerComposeCmdInDirs(dirs, "down")
}

func dockerPull() {
	dirs, err := findEnabledDirs(".", false)
	if err != nil {
		fmt.Printf("Error scanning directories: %v\n", err)
		os.Exit(1)
	}
	runDockerComposeCmdInDirs(dirs, "pull")
}

func dockerRestart() {
	dockerDown()
	dockerUp()
}

func dockerStart() {
	dirs, err := findEnabledDirs(".", false)
	if err != nil {
		fmt.Printf("Error scanning directories: %v\n", err)
		os.Exit(1)
	}
	runDockerComposeCmdInDirs(dirs, "start")
}

func dockerStop() {
	dirs, err := findEnabledDirs(".", true)
	if err != nil {
		fmt.Printf("Error scanning directories: %v\n", err)
		os.Exit(1)
	}
	runDockerComposeCmdInDirs(dirs, "stop")
}
