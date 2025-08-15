// cmd/configure.go
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/techspeque/codelm/internal/config"
	"gopkg.in/yaml.v3"
)

// A temporary struct to build the config before writing it to a file.
type configScaffold struct {
	ApiKeys struct {
		OpenAI    string `yaml:"openai"`
		Anthropic string `yaml:"anthropic"`
	} `yaml:"api_keys"`
	Shell struct {
		ImportUserSettings bool `yaml:"import_user_settings"`
	} `yaml:"shell"`
}

// Main 'configure' command
var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Create or show the codelm configuration file",
	Long:  "Run without a subcommand to create a new configuration. Use 'show' to display the current config.",
	Run: func(cmd *cobra.Command, args []string) {
		// This is the interactive setup from before
		reader := bufio.NewReader(os.Stdin)
		scaffold := configScaffold{}

		fmt.Println("--- codelm Configuration ---")

		fmt.Print("Enter your OpenAI API key (leave blank to skip): ")
		scaffold.ApiKeys.OpenAI, _ = reader.ReadString('\n')
		scaffold.ApiKeys.OpenAI = strings.TrimSpace(scaffold.ApiKeys.OpenAI)

		fmt.Print("Enter your Anthropic API key (leave blank to skip): ")
		scaffold.ApiKeys.Anthropic, _ = reader.ReadString('\n')
		scaffold.ApiKeys.Anthropic = strings.TrimSpace(scaffold.ApiKeys.Anthropic)

		fmt.Print("Import your personal shell settings (~/.zshrc, etc.)? (y/N): ")
		importSettings, _ := reader.ReadString('\n')
		if strings.TrimSpace(strings.ToLower(importSettings)) == "y" {
			scaffold.Shell.ImportUserSettings = true
		}

		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("Fatal: Could not get user home directory.", err)
			os.Exit(1)
		}
		configPath := filepath.Join(home, ".config", "codelm")
		configFilePath := filepath.Join(configPath, "config.yaml")

		data, err := yaml.Marshal(&scaffold)
		if err != nil {
			fmt.Println("Fatal: Could not format config data.", err)
			os.Exit(1)
		}

		if err := os.MkdirAll(configPath, 0755); err != nil {
			fmt.Println("Fatal: Could not create config directory.", err)
			os.Exit(1)
		}
		if err := os.WriteFile(configFilePath, data, 0600); err != nil {
			fmt.Println("Fatal: Could not write config file.", err)
			os.Exit(1)
		}

		fmt.Printf("\n✅ Configuration saved successfully to %s\n", configFilePath)
	},
}

// New 'show' subcommand
var showConfigCmd = &cobra.Command{
	Use:   "show",
	Short: "Show the current configuration (with redacted keys)",
	Run: func(cmd *cobra.Command, args []string) {
		configPath := viper.ConfigFileUsed()
		if configPath == "" {
			fmt.Println("Configuration file not found. Please run 'codelm configure' first.")
			os.Exit(1)
		}

		fmt.Printf("Displaying configuration from: %s\n\n", configPath)
		fmt.Println("api_keys:")
		fmt.Printf("  openai: %s\n", redactKey(config.Cfg.ApiKeys.OpenAI))
		fmt.Printf("  anthropic: %s\n", redactKey(config.Cfg.ApiKeys.Anthropic))
		fmt.Println("shell:")
		fmt.Printf("  import_user_settings: %t\n", config.Cfg.Shell.ImportUserSettings)
	},
}

// redactKey checks if a key is set and returns a redacted status.
func redactKey(key string) string {
	if key != "" {
		return "<set>"
	}
	return "<not set>"
}

func init() {
	// Add 'show' as a subcommand of 'configure'
	configureCmd.AddCommand(showConfigCmd)
	rootCmd.AddCommand(configureCmd)
}