package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

// Config struct matches the structure of your config.yaml
type Config struct {
	LastProjectPath string `mapstructure:"last_project_path" yaml:"last_project_path"` // Added this field
	ApiKeys         struct {
		OpenAI    string `mapstructure:"openai" yaml:"openai"`
		Anthropic string `mapstructure:"anthropic" yaml:"anthropic"`
	} `mapstructure:"api_keys" yaml:"api_keys"`
	Shell struct {
		ImportUserSettings bool `mapstructure:"import_user_settings" yaml:"import_user_settings"`
	} `mapstructure:"shell" yaml:"shell"`
}

var Cfg *Config

// SaveConfig saves the current configuration back to the config file.
func SaveConfig() error {
	if Cfg == nil {
		return fmt.Errorf("config is not initialized")
	}

	configFilePath := viper.ConfigFileUsed()
	if configFilePath == "" {
		// If the file doesn't exist, create it in the default location
		home, err := os.UserHomeDir()
		if err != nil {
			return err
		}
		configPath := filepath.Join(home, ".config", "codelm")
		configFilePath = filepath.Join(configPath, "config.yaml")
		if err := os.MkdirAll(configPath, 0755); err != nil {
			return err
		}
	}

	// Marshal the Cfg struct back to YAML
	data, err := yaml.Marshal(&Cfg)
	if err != nil {
		return fmt.Errorf("could not format config data: %w", err)
	}

	// Write the file with secure permissions
	if err := os.WriteFile(configFilePath, data, 0600); err != nil {
		return fmt.Errorf("could not write config file: %w", err)
	}

	return nil
}

func InitConfig() {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error finding home directory:", err)
		os.Exit(1)
	}

	configPath := filepath.Join(home, ".config", "codelm")
	viper.AddConfigPath(configPath)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; initialize an empty config
			Cfg = &Config{}
		} else {
			fmt.Println("Error reading config file:", err)
			os.Exit(1)
		}
	}

	if Cfg == nil {
		err = viper.Unmarshal(&Cfg)
		if err != nil {
			fmt.Println("Unable to decode into struct:", err)
			os.Exit(1)
		}
	}
}