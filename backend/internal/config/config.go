package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	BackendAPI      string
	DBContainerName string
	DBName          string
	DBUser          string
	DBPassword      string
}

func getEnv(env, defaultValue string) string {
	if value := os.Getenv(env); value != "" {
		return value
	}
	return defaultValue
}

func NewConfig() *Config {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("Warning: .env file not found, using environment variables")
	}

	return &Config{
		BackendAPI:      getEnv("BACKEND_API", ":8080"),
		DBContainerName: getEnv("DB_CONTAINER_NAME", "db"),
		DBName:          getEnv("DB_NAME", "todo"),
		DBUser:          getEnv("DB_USER", "postgres"),
		DBPassword:      getEnv("DB_PASSWORD", "postgres"),
	}
}
