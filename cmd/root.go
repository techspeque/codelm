package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/techspeque/codelm/internal/config"
)

var rootCmd = &cobra.Command{
	Use:   "codelm",
	Short: "codelm creates isolated, containerized environments for AI-assisted coding.",
	Long:  `A CLI utility that manages Docker containers with pre-installed AI tools to provide a consistent and secure development workspace.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your command '%s'", err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(config.InitConfig)
}