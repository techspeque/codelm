// cmd/stop.go
package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/techspeque/codelm/internal/config"
	"github.com/techspeque/codelm/internal/docker"
)

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop and remove the codelm container session",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Stopping codelm session...")

		// --- NEW --- Inspect container to get project path before stopping
		state, err := docker.GetContainerState()
		if err == nil { // If container exists
			for _, mount := range state.Mounts {
				if strings.HasSuffix(mount, ":/workspace") {
					hostPath := strings.TrimSuffix(mount, ":/workspace")
					config.Cfg.LastProjectPath = hostPath
					if err := config.SaveConfig(); err != nil {
						fmt.Println("Warning: Could not save last project path.", err)
					}
					break
				}
			}
		}
		// --- END NEW ---

		if err := docker.StopAndRemoveContainer(); err != nil {
			fmt.Println("An error occurred while stopping the session.")
		} else {
			fmt.Println("Session stopped successfully.")
		}
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
}