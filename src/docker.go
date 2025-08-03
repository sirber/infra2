package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func findEnabledDirs(root string) ([]string, error) {
	var dirs []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			// Skip directories/files we can't access
			if os.IsPermission(err) {
				return filepath.SkipDir
			}
			return nil
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
		fmt.Printf("Running 'docker compose %v' in %s\n", args, dir)
		if err := cmd.Run(); err != nil {
			fmt.Printf("Error running docker compose in %s: %v\n", dir, err)
		}
	}
}

func dockerUp() {
	dirs, err := findEnabledDirs(".")
	if err != nil {
		fmt.Printf("Error scanning directories: %v\n", err)
		os.Exit(1)
	}
	runDockerComposeCmdInDirs(dirs, "up", "-d")
}

func dockerDown() {
	dirs, err := findEnabledDirs(".")
	if err != nil {
		fmt.Printf("Error scanning directories: %v\n", err)
		os.Exit(1)
	}
	runDockerComposeCmdInDirs(dirs, "down")
}

func dockerPull() {
	dirs, err := findEnabledDirs(".")
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
	dirs, err := findEnabledDirs(".")
	if err != nil {
		fmt.Printf("Error scanning directories: %v\n", err)
		os.Exit(1)
	}
	runDockerComposeCmdInDirs(dirs, "start")
}

func dockerStop() {
	dirs, err := findEnabledDirs(".")
	if err != nil {
		fmt.Printf("Error scanning directories: %v\n", err)
		os.Exit(1)
	}
	runDockerComposeCmdInDirs(dirs, "stop")
}

func getComposeFilesFromEnabledDirs() ([]string, error) {
	dirs, err := findEnabledDirs(".")
	if err != nil {
		return nil, err
	}
	var composeFiles []string
	for _, dir := range dirs {
		composePath := filepath.Join(dir, "docker-compose.yml")
		if _, err := os.Stat(composePath); err == nil {
			composeFiles = append(composeFiles, composePath)
		} else if os.IsPermission(err) {
			// Skip files we can't access
			continue
		}
	}
	return composeFiles, nil
}

func dockerStatus() {
	composeFiles, err := getComposeFilesFromEnabledDirs()
	if err != nil {
		fmt.Printf("Error scanning directories: %v\n", err)
		os.Exit(1)
	}
	if len(composeFiles) == 0 {
		fmt.Println("No docker-compose.yml files found in enabled folders.")
		return
	}
	args := []string{"compose"}
	for _, f := range composeFiles {
		args = append(args, "-f", f)
	}
	args = append(args, "ps")
	cmd := exec.Command("docker", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	fmt.Printf("Running 'docker %v'\n", args)
	if err := cmd.Run(); err != nil {
		fmt.Printf("Error running docker compose ps: %v\n", err)
	}
}

func dockerLogs() {
	composeFiles, err := getComposeFilesFromEnabledDirs()
	if err != nil {
		fmt.Printf("Error scanning directories: %v\n", err)
		os.Exit(1)
	}
	if len(composeFiles) == 0 {
		fmt.Println("No docker-compose.yml files found in enabled folders.")
		return
	}
	args := []string{"compose"}
	for _, f := range composeFiles {
		args = append(args, "-f", f)
	}
	args = append(args, "logs", "-f") // follow logs
	cmd := exec.Command("docker", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	fmt.Printf("Running 'docker %v'\n", args)
	if err := cmd.Run(); err != nil {
		fmt.Printf("Error running docker compose logs: %v\n", err)
	}
}
