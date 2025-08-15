package cmd

import (
	"fmt"
	"os"
	"os/user" // --- NEW --- Import the 'user' package
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/techspeque/codelm/internal/config"
	"github.com/techspeque/codelm/internal/docker"
)

var startCmd = &cobra.Command{
	Use:   "start [path-to-project]",
	Short: "Starts the codelm session, using the last project if no path is given",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := docker.EnsureImageExists(); err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		if config.Cfg == nil {
			fmt.Println("Configuration not loaded. Please run 'codelm configure' first.")
			os.Exit(1)
		}

		var projectPath string
		if len(args) > 0 {
			projectPath = args[0]
			config.Cfg.LastProjectPath = projectPath
			if err := config.SaveConfig(); err != nil {
				fmt.Println("Warning: Could not save the new project path to config.", err)
			}
		} else {
			if config.Cfg.LastProjectPath == "" {
				fmt.Println("Error: No project path specified and no saved project found.")
				os.Exit(1)
			}
			projectPath = config.Cfg.LastProjectPath
		}

		absPath, err := filepath.Abs(projectPath)
		if err != nil {
			fmt.Println("Error: Could not determine absolute path for your project.", err)
			os.Exit(1)
		}
		fmt.Printf("Starting codelm session for project: %s\n", absPath)
		docker.StopAndRemoveContainer()

		// --- NEW --- Get current user's UID and GID
		currentUser, err := user.Current()
		if err != nil {
			fmt.Println("Warning: Could not determine current user. Files created in the container may have incorrect permissions.", err)
			os.Exit(1)
		}
		userAndGroup := fmt.Sprintf("%s:%s", currentUser.Uid, currentUser.Gid)
		// --- END NEW ---

		dockerArgs := []string{
			"run",
			"-itd",
			"--name", docker.ContainerName,
			"--user", userAndGroup, // --- NEW --- Add the user flag
			"-v", fmt.Sprintf("%s:/workspace", absPath),
		}

		// (The rest of the file is the same)
		if config.Cfg.ApiKeys.OpenAI != "" {
			dockerArgs = append(dockerArgs, "-e", fmt.Sprintf("OPENAI_API_KEY=%s", config.Cfg.ApiKeys.OpenAI))
		}
		if config.Cfg.ApiKeys.Anthropic != "" {
			dockerArgs = append(dockerArgs, "-e", fmt.Sprintf("ANTHROPIC_API_KEY=%s", config.Cfg.ApiKeys.Anthropic))
		}
		if config.Cfg.Shell.ImportUserSettings {
			fmt.Println("Importing user shell settings...")
			home, _ := os.UserHomeDir()
			dockerArgs = append(dockerArgs, "-v", fmt.Sprintf("%s:/home/devuser/.zshrc_user:ro", filepath.Join(home, ".zshrc")))
			dockerArgs = append(dockerArgs, "-v", fmt.Sprintf("%s:/home/devuser/.zsh_history", filepath.Join(home, ".zsh_history")))
			dockerArgs = append(dockerArgs, "-v", fmt.Sprintf("%s:/home/devuser/.gitconfig:ro", filepath.Join(home, ".gitconfig")))
		}
		dockerArgs = append(dockerArgs, docker.ImageName)

		fmt.Println("Starting container...")
		if err := docker.RunCommand("docker", dockerArgs...); err != nil {
			fmt.Println("Error starting Docker container. Is the Docker daemon running?", err)
			os.Exit(1)
		}

		fmt.Println("\n✅ Session started successfully!")
		fmt.Println("Run 'codelm shell' to get inside the container.")
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}