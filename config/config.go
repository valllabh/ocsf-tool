package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// initConfig reads in config file and ENV variables if set.
func InitConfig(configFile string) error {
	viper.SetConfigType("yaml")
	viper.SetConfigFile(configFile)

	return viper.ReadInConfig()
}

// Write config file to disk
func WriteConfig() {
	err := viper.WriteConfigAs("config.yaml")

	// Handle errors writing the config file
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	println("Config saved.")
}
