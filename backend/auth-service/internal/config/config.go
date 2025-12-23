package config

import (
	"os"
)

type Config struct {
	Port        string
	Environment string
	JWTSecret   string
	DatabaseURL string
}

func Load() *Config {
	return &Config{
		Port:        getEnv("PORT", "3000"),
		Environment: getEnv("ENVIRINMENRT", "development"),
		JWTSecret:   getEnv("JWT_SECRET", "auth-service-secret"),
		DatabaseURL: getEnv("DATABASE_URL", "postgres://admin:password@localhost:5432/auth_db"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
