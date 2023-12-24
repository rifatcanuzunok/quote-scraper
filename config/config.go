package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config holds the configuration settings
type Config struct {
	DBUsername string
	DBPassword string
	DBHost     string
	DBPort     int
	DBName     string
}

// LoadConfig loads configuration settings from environment variables or config files
func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	var config Config

	// Load configuration from environment variables or config file
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Config file not found. Reading from environment variables.")
		config.DBUsername = viper.GetString("DB_USERNAME")
		config.DBPassword = viper.GetString("DB_PASSWORD")
		config.DBHost = viper.GetString("DB_HOST")
		config.DBPort = viper.GetInt("DB_PORT")
		config.DBName = viper.GetString("DB_NAME")
	} else {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
		if err := viper.Unmarshal(&config); err != nil {
			return nil, err
		}
	}

	return &config, nil
}
