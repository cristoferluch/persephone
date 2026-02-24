package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Host     string
	User     string
	Password string
	DataBase string
	Port     int
	SSLMode  string
}

func Load() (*Config, error) {

	viper.SetConfigName("settings")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("error reading configs file, %w", err)
	}

	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, fmt.Errorf("unable to decode into struct, %w", err)
	}
	return &config, nil
}
