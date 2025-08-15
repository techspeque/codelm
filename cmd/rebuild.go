// cmd/rebuild.go
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/techspeque/codelm/internal/docker"
)

var rebuildCmd = &cobra.Command{
	Use:   "rebuild",
	Short: "Forcefully rebuilds the codelm Docker image",
	Run: func(cmd *cobra.Command, args []string) {
		if err := docker.RemoveImage(); err != nil {
			// Don't exit on error, as the image might not have existed
			fmt.Println("Could not remove old image (it may not have existed). Continuing...")
		}

		if err := docker.BuildImage(); err != nil {
			fmt.Println("An error occurred during the build process.", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(rebuildCmd)
}