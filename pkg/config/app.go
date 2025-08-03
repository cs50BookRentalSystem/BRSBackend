package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type AppConfig struct {
	Server    ServerConfig    `mapstructure:"server"`
	Database  DatabaseConfig  `mapstructure:"database"`
	Librarian LibrarianConfig `mapstructure:"librarian"`
	Rent      RentalConfig    `mapstructure:"rent"`
}

type ServerConfig struct {
	Port string `mapstructure:"port"`
	Env  string `mapstructure:"env"`
}

type DatabaseConfig struct {
	DSN string `mapstructure:"dsn"`
}

type LibrarianConfig struct {
	User string `mapstructure:"user"`
	Pass string `mapstructure:"pass"`
}

type RentalConfig struct {
	RentalDays int `mapstructure:"rental_days"`
}

func LoadConfig(path string) (*AppConfig, error) {
	viper.SetConfigFile(path)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var config AppConfig
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &config, nil
}
