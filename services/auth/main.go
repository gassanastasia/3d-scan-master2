package maim

import (
	"auth-service/internal/config"
	"auth-service/internal/database"
	"auth-service/internal/handlers"
	"auth-service/internal/server"
	"log"
)

func main() {
	cfg := config.Load()

	db, err := database.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	authHandler := handlers.NewAuthHandler(db, cfg.JWTSecret)

	srv := server.NewServer(cfg, authHandler)
}