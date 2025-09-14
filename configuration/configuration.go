package configuration

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/zelcion/focusmode/constants"
)

func getDefaultConfigPath() string {
	homeDir, err := os.UserHomeDir()

	if err != nil {
		var errorMessage = fmt.Sprintf("Failed to get user home directory: %v", err)
		panic(errorMessage)
	}

	return filepath.Join(homeDir, ".config", constants.CONFIG_DIR_NAME, constants.CONFIG_FILE_NAME)
}

func getDefaultConfigDir() string {
	configPath := getDefaultConfigPath()
	return filepath.Dir(configPath)
}

func ensureConfigDirExists() error {
	configDir := getDefaultConfigDir()

	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		mkdirErr := os.MkdirAll(configDir, os.ModePerm)
		if mkdirErr != nil {
			return fmt.Errorf("failed to create config directory: %v", mkdirErr)
		}
	}

	return nil
}

func InitializeConfig() (*Configuration, error) {
	if err := ensureConfigDirExists(); err != nil {
		return nil, err
	}

	configPath := getDefaultConfigPath()
	var config *Configuration

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		config = NewEmptyConfiguration()
		config.Save(configPath)
		return config, nil
	}

	return LoadConfiguration(configPath)
}
