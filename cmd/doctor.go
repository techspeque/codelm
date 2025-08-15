// cmd/doctor.go
package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/techspeque/codelm/internal/config"
)

var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Check your environment for potential problems",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("🩺 Running codelm doctor...")
		var checksFailed bool

		fmt.Print("   - Checking for Docker...")
		_, err := exec.LookPath("docker")
		if err != nil {
			fmt.Println(" ✗ Not Found")
			fmt.Println("     ↪ Docker is not installed or not in your PATH.")
			fmt.Println("       Please install it from https://www.docker.com/products/docker-desktop/")
			os.Exit(1) // This is a fatal error
		}
		fmt.Println(" ✓ Installed")

		fmt.Print("   - Checking Docker daemon...")
		if err := exec.Command("docker", "info").Run(); err != nil {
			fmt.Println(" ✗ Not running")
			fmt.Println("     ↪ Please start the Docker daemon and try again.")
			checksFailed = true
		} else {
			fmt.Println(" ✓ Running")
		}

		fmt.Print("   - Checking for config file...")
		configFilePath := viper.ConfigFileUsed()
		if configFilePath == "" {
			fmt.Println(" ✗ Not found")
			fmt.Println("     ↪ Please run 'codelm configure' to create it.")
			checksFailed = true
		} else {
			fmt.Printf(" ✓ Found at %s\n", configFilePath)
			fmt.Print("   - Checking API key configuration...")
			if config.Cfg.ApiKeys.OpenAI == "" && config.Cfg.ApiKeys.Anthropic == "" {
				fmt.Println(" 🟡 None configured")
			} else {
				fmt.Println(" ✓ At least one key is configured.")
			}
		}

		fmt.Println()
		if checksFailed {
			fmt.Println("Doctor found one or more critical issues to resolve.")
			os.Exit(1)
		} else {
			fmt.Println("✅ Your environment looks good!")
		}
	},
}

func init() {
	rootCmd.AddCommand(doctorCmd)
}