package helpers

import (
	"github.com/joho/godotenv"
	"os"
)

func Getenv(key string) string {
	_ = godotenv.Load()
	return os.Getenv(key)
}