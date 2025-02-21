package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBPath      string
	ServerPort  string
	Environment string
}

func LoadConfig() *Config {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}

	return &Config{
		DBPath:      getEnvOrDefault("DB_PATH", "./todos.db"),
		ServerPort:  getEnvOrDefault("SERVER_PORT", "8080"),
		Environment: getEnvOrDefault("ENVIRONMENT", "development"),
	}
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
