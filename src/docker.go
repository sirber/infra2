package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func findEnabledDirs(root string) ([]string, error) {
	var dirs []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
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
		log.Printf("Running 'docker compose %v' in %s", args, dir)
		if err := cmd.Run(); err != nil {
			log.Printf("Error running docker compose in %s: %v", dir, err)
		}
	}
}

func dockerUp() {
	dirs, err := findEnabledDirs(".")
	if err != nil {
		log.Fatalf("Error scanning directories: %v", err)
	}
	runDockerComposeCmdInDirs(dirs, "up", "-d")
}

func dockerDown() {
	dirs, err := findEnabledDirs(".")
	if err != nil {
		log.Fatalf("Error scanning directories: %v", err)
	}
	runDockerComposeCmdInDirs(dirs, "down")
}

func dockerPull() {
	dirs, err := findEnabledDirs(".")
	if err != nil {
		log.Fatalf("Error scanning directories: %v", err)
	}
	runDockerComposeCmdInDirs(dirs, "pull")
}

func dockerRestart() {
	dockerDown()
	dockerUp()
}

func dockerStart() {
	dirs, err := findEnabledDirs(".")
	if err != nil {
		log.Fatalf("Error scanning directories: %v", err)
	}
	runDockerComposeCmdInDirs(dirs, "start")
}

func dockerStop() {
	dirs, err := findEnabledDirs(".")
	if err != nil {
		log.Fatalf("Error scanning directories: %v", err)
	}
	runDockerComposeCmdInDirs(dirs, "stop")
}
