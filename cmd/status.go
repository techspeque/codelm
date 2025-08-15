// cmd/status.go
package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/techspeque/codelm/internal/docker"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show the status of the codelm container session",
	Run: func(cmd *cobra.Command, args []string) {
		state, err := docker.GetContainerState()
		if err != nil {
			fmt.Println("⚪️ codelm session is not running.")
			fmt.Printf("   └── Debug: %v\n", err)
			return
		}

		if state.IsRunning {
			fmt.Println("🟢 codelm session is running.")
		} else {
			fmt.Printf("🟡 codelm session is stopped (Status: %s).\n", state.Status)
		}

		fmt.Printf("   - Image: %s\n", state.Image)

		for _, mount := range state.Mounts {
			if strings.HasSuffix(mount, ":/workspace") {
				hostPath := strings.TrimSuffix(mount, ":/workspace")
				fmt.Printf("   - Project: %s\n", hostPath)
				break
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}