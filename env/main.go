// handles loading and validation of environment variable

package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PORT        string
	ENVIRONMENT string
}

var DefaultConfig Config

func Load() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatalln("$PORT must be set")
	}

	environment := os.Getenv("ENVIRONMENT")
	if environment == "" {
		log.Fatalln("$ENVIRONMENT must be set")
	}

	DefaultConfig = Config{
		PORT:        port,
		ENVIRONMENT: environment,
	}

	log.Println("Loaded .env!")
}
