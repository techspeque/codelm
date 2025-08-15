package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/techspeque/codelm/internal/docker"
)

var shellCmd = &cobra.Command{
	Use:   "shell",
	Short: "Attach a shell to the running codelm container",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Attaching to codelm session...")
		err := docker.RunCommand("docker", "exec", "-it", docker.ContainerName, "/bin/zsh")
		if err != nil {
			fmt.Printf("Error attaching to container. Is a session running? (start with 'codelm start')\n")
		}
	},
}

func init() {
	rootCmd.AddCommand(shellCmd)
}