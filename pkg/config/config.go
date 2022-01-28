package config

import (
	// "log"
	"os"
	"path/filepath"
	
	"github.com/spf13/viper"
)

const (
	configPath          string = "./"
	configFileType      string = "yaml"
	configName          string = "config"
	configFileExtension string = ".yml"
	productionEnv       string = "Production"
)

var Config *Configuration = New()

func New() (*Configuration) {

	viper.SetConfigType(configFileType)
	viper.SetConfigName(configName)
	viper.AddConfigPath(configPath)

	viper.BindEnv("environment.env", "B_ENV")
	viper.SetDefault("environment.env", "Development")
	viper.SetDefault("log.path", "onent-devops.log")
	viper.SetDefault("server.port", "8416")
	viper.SetDefault("server.timeout", 50)
	configFilePath := filepath.Join(configPath, configName) + configFileExtension
	if err := readConfiguration(configFilePath); err != nil {
		return nil
	}

	viper.AutomaticEnv()
	var cfg *Configuration
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil
	}

	viper.WatchConfig()

	return cfg
}

// read configuration from file
func readConfiguration(filePath string) error {
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		// if file does not exist, simply create one
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			os.Create(filePath)
		} else {
			return err
		}
		// let's write defaults
		if err := viper.WriteConfig(); err != nil {
			return err
		}
	}
	return nil
}
