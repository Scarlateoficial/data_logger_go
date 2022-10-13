package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvMongoURI() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("MONGOURI")
}

type Config struct {
	mqtt struct {
		host     string
		port     int
		user     string
		password string
		name     string
	}
	server struct {
		host string
		port int
	}
}

func getConfig() Config {
	return Config{}
}
