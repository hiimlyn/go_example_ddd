package config

import (
	"os"
)

type Config struct {
	Port  string
	DBDns string
}

func New() *Config {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	return &Config{
		Port:  os.Getenv("PORT"),
		DBDns: os.Getenv("DB_DNS"),
	}
}
