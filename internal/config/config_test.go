// internal/config/config_test.go
package config

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

// TestSaveConfig checks if the configuration can be written to a file correctly.
func TestSaveConfig(t *testing.T) {
	// Create a temporary directory for our test config file.
	// This is automatically cleaned up when the test finishes.
	tempDir := t.TempDir()
	configFilePath := filepath.Join(tempDir, "config.yaml")

	// Set up Viper to use our temporary file.
	viper.SetConfigFile(configFilePath)

	// Define the configuration we want to save.
	originalCfg := &Config{
		LastProjectPath: "/path/to/my/project",
		ApiKeys: struct {
			OpenAI    string `mapstructure:"openai" yaml:"openai"`
			Anthropic string `mapstructure:"anthropic" yaml:"anthropic"`
		}{
			OpenAI: "test-openai-key",
		},
		Shell: struct {
			ImportUserSettings bool `mapstructure:"import_user_settings" yaml:"import_user_settings"`
		}{
			ImportUserSettings: true,
		},
	}

	// Assign our test config to the global Cfg variable, as the function uses it.
	Cfg = originalCfg

	// Run the function we want to test.
	err := SaveConfig()
	if err != nil {
		t.Fatalf("SaveConfig() returned an unexpected error: %v", err)
	}

	// Now, read the file back to verify its contents.
	data, err := os.ReadFile(configFilePath)
	if err != nil {
		t.Fatalf("Failed to read back config file: %v", err)
	}

	// Unmarshal the YAML data from the file into a new struct.
	var savedCfg Config
	err = yaml.Unmarshal(data, &savedCfg)
	if err != nil {
		t.Fatalf("Failed to unmarshal saved config: %v", err)
	}

	// Compare the saved config with the original.
	if savedCfg.LastProjectPath != originalCfg.LastProjectPath {
		t.Errorf("Expected LastProjectPath to be %q, but got %q", originalCfg.LastProjectPath, savedCfg.LastProjectPath)
	}
	if savedCfg.ApiKeys.OpenAI != originalCfg.ApiKeys.OpenAI {
		t.Errorf("Expected OpenAI API key to be %q, but got %q", originalCfg.ApiKeys.OpenAI, savedCfg.ApiKeys.OpenAI)
	}
	if savedCfg.Shell.ImportUserSettings != originalCfg.Shell.ImportUserSettings {
		t.Errorf("Expected ImportUserSettings to be %t, but got %t", originalCfg.Shell.ImportUserSettings, savedCfg.Shell.ImportUserSettings)
	}
}