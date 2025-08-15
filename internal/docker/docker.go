// internal/docker/docker.go
package docker

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
)

const (
	ContainerName = "codelm_session"
	ImageName     = "codelm:latest"
)

// ContainerState holds the information we want from 'docker inspect'
type ContainerState struct {
	Name      string
	Status    string
	Image     string
	Mounts    []string
	IsRunning bool
}

// BuildImage builds the codelm Docker image
func BuildImage() error {
	fmt.Println("Building Docker image 'codelm:latest'...")
	fmt.Println("This may take a few minutes.")
	buildCmd := exec.Command("docker", "build", "-t", ImageName, "build/")
	buildCmd.Stdout = os.Stdout
	buildCmd.Stderr = os.Stderr
	if err := buildCmd.Run(); err != nil {
		return fmt.Errorf("failed to build docker image: %w", err)
	}
	fmt.Println("✅ Docker image built successfully.")
	return nil
}

// RemoveImage forcefully removes the codelm Docker image
func RemoveImage() error {
	fmt.Println("Removing existing Docker image...")
	// Use -f to force remove even if containers are using it
	cmd := exec.Command("docker", "rmi", "-f", ImageName)
	cmd.Stdout = io.Discard
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// EnsureImageExists checks if the codelm Docker image exists, and if not, builds it.
func EnsureImageExists() error {
	cmd := exec.Command("docker", "image", "inspect", ImageName)
	cmd.Stdout = io.Discard
	cmd.Stderr = io.Discard
	if cmd.Run() != nil {
		fmt.Println("Docker image 'codelm:latest' not found.")
		return BuildImage() // Call the new build function
	}
	return nil
}

// GetContainerState inspects the container and returns its state.
func GetContainerState() (*ContainerState, error) {
	const format = `{"Name": "{{.Name}}", "Status": "{{.State.Status}}", "Image": "{{.Config.Image}}", "Mounts": {{json .HostConfig.Binds}}, "IsRunning": {{.State.Running}}}`

	cmd := exec.Command("docker", "inspect", "--format", format, ContainerName)

	output, err := cmd.Output()
	if err != nil {
		// This error typically means the container doesn't exist
		return nil, err
	}

	var state ContainerState
	if err := json.Unmarshal(output, &state); err != nil {
		// This would be an unexpected error with our JSON format
		return nil, fmt.Errorf("error parsing docker inspect output: %w", err)
	}

	return &state, nil
}

// RunCommand executes a system command and prints its output.
func RunCommand(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin // For interactive commands like 'shell'
	return cmd.Run()
}

// StopAndRemoveContainer stops and removes the codelm container if it exists.
func StopAndRemoveContainer() error {
	// The 'docker stop' command will error if the container doesn't exist, which is fine.
	// We ignore the error for stop and proceed to rm.
	exec.Command("docker", "stop", ContainerName).Run()

	fmt.Println("Removing existing container...")
	err := exec.Command("docker", "rm", ContainerName).Run()
	if err != nil {
		// Suppress "No such container" errors to keep the output clean.
	}
	return nil
}