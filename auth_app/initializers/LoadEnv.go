package initializers

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func LoadEnv() {
	env := os.Getenv("ENV_FILE")

	if "" == env {
		env = "dev"
	}

	err := godotenv.Load(".env." + env)

	if err != nil {
		log.Fatal(err)
	}
}
