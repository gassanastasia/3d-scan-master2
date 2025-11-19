package config

import "os"

type Config struct {
	Port       string
	Enviroment string
	JWTSecret  string
	ClickHouseURL string
	RedisURL string
	KafkaBrokers string
	KlipperAPI string
}

func Load() *Config {
	return &Config{
		Port:       getEnv("PORT", "8080"),
		Enviroment: getEnv("ENVIROMENT", "development"),
		JWTSecret:  getEnv("JWT_SECRET", "dev-secret-key"),
		ClickHouseURL:  getEnv("CLICKHOUSE_URL", "tcp://localhost:9000"),
		RedisURL:  getEnv("REDIS_URL", "localhost:6379"),
		KafkaBrokers:  getEnv("KAFKA_BROKERS", "localhost:9092"),
		KlipperAPI:  getEnv("KLIPPER_API", "http://localhost:7125"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
