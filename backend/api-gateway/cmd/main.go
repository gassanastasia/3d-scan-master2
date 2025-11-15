package main

import (
    "log"
    "api-gateway/internal/server"
    "api-gateway/internal/config"
)

func main() {
    cfg := config.Load()
    srv := server.NewServer(cfg)
    
    log.Printf("API Gateway starting on :%s", cfg.Port)
    if err := srv.Run(); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}