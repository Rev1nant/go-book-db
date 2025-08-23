package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB DBConfig
}

type DBConfig struct {
	Dsn string
}

func LoadConfig(path string) *Config {
	if err := godotenv.Load(path); err != nil {
		log.Println("not .env file found")
	}
	return &Config{
		DB: DBConfig{
			Dsn: os.Getenv("DSN"),
		},
	}
}
