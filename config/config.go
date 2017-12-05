package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// LoadConfig loads the configuration in the viper.
func LoadConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath("config/")

	err := viper.ReadInConfig()
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s ", err))
	}

}
