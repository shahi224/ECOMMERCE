package config

import (
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type Config struct {
	ServerAddress string `mapstructure:"SERVER_ADDRESS" validate:"required"`
	DBHost        string `mapstructure:"DB_HOST" validate:"required"`
	DBName        string `mapstructure:"DB_NAME" validate:"required"`
	DBUser        string `mapstructure:"DB_USER" validate:"required"`
	DBPort        string `mapstructure:"DB_PORT" validate:"required"`
	DBPassword    string `mapstructure:"DB_PASSWORD" validate:"required"`
	Key           string `mapstructure:"KEY" validate:"required"`
}

func LoadConfig() (Config, error) {
	var config Config
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Faled to load config file: %v", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Failed to unmarshal config: %v", err)
	}
	if err := validator.New().Struct(&config); err != nil {
		log.Fatalf("Invalid config structure: %v", err)
	}
	return config, nil
}
