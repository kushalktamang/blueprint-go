package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PORT string
	ENV  string
}

func MustLoad() Config {
	godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		panic("port is required")
	}

	env := os.Getenv("ENV")
	if port == "" {
		panic("environment is required")
	}

	return Config{
		PORT: port,
		ENV:  env,
	}

}
