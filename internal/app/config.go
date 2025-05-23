package app

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func LoadEnvironment() {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("error getting current directory: %s", err.Error())
	}

	for {
		envPath := filepath.Join(currentDir, ".env")

		if _, err := os.Stat(envPath); err == nil {
			if err := godotenv.Load(envPath); err != nil {
				log.Fatalf("error loading env variables: %s", err.Error())
			}
			return
		}

		parentDir := filepath.Dir(currentDir)
		if parentDir == currentDir {
			return
		}
		currentDir = parentDir
	}
}
