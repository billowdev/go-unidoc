package configs

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

func findConfigFile() (string, error) {
	// Get the current working directory
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
		// return "", err
	}

	// Search for the config file recursively upwards
	for {
		configPath := filepath.Join(cwd, "config.yaml")
		if _, err := os.Stat(configPath); err == nil {
			return configPath, nil // Configuration file found
		}

		// Move one directory up
		parentDir := filepath.Dir(cwd)
		if parentDir == cwd {
			// Reached the root directory
			break
		}
		cwd = parentDir
	}

	return "", fmt.Errorf("config file not found")
}
func InitConfig() error {
	// Find the configuration file
	configPath, err := findConfigFile()
	if err != nil {
		return err
	}
	// Set the config file
	viper.SetConfigFile(configPath)
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	// Read the config file
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	// fmt.Println("Config loaded successfully")
	return nil
}

func init() {
	InitConfig()
}
