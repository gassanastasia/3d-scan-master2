package config

import "os"

type Config struct {
	Port       string
	Enviroment string
	JWTSecret  string
}

func Load() *Config {
	return &Config{
		Port:       getEnv("PORT", "8080"),
		Enviroment: getEnv("ENVIROMENT", "development"),
		JWTSecret:  getEnv("JWT_SECRET", "dev-secret-key"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
