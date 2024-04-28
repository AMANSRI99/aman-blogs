package config

import (
	"log"
	"os"

	"github.com/AMANSRI99/aman-blogs/utils/errs"
	"github.com/joho/godotenv"
)

const (
	PORT     = "APP_PORT"
	ENV_FILE = ".env"
)

func init() {
	if err := godotenv.Load(ENV_FILE); err != nil {
		log.Fatal(errs.Wrap(err))
	}
}

func Get(key string) string {
	return os.Getenv(key)
}
