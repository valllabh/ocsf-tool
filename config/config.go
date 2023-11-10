package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func getConfigFilePath() string {
	return "config.yaml"
}

// initConfig reads in config file and ENV variables if set.
func InitConfig() {

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	// Read in environment variables that match
	err := viper.ReadInConfig()

	// Handle errors reading the config file
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; create it with default values
			viper.WriteConfigAs(getConfigFilePath()) // handle error
		} else {
			// Config file was found but another error was produced
			panic(fmt.Errorf("fatal error config file: %s", err))
		}
	}
}

// Write config file to disk
func WriteConfig() {
	err := viper.WriteConfigAs(getConfigFilePath())

	// Handle errors writing the config file
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	println("Config saved.")
}
