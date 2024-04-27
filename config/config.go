package config

import "os"

const (
	PORT = "APP_PORT"
)

func Get(key string) string {
	return os.Getenv(key)
}
