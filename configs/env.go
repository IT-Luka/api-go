package configs

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func EnvMongoURI() string {
	ginEnv := os.Getenv("GIN_MODE")

	if ginEnv != "release" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
	return os.Getenv("MONGO_URI")
}
