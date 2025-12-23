package main

import (
	"api-gateway/internal/config"
	"api-gateway/internal/server"
	"log"

	"github.com/golang-jwt/jwt/request"
	"github.com/prometheus/client_golang/prometheus"
)

func main() {
    cfg := config.Load()
    srv := server.NewServer(cfg)

    initMetrics()
    
    log.Printf("API Gateway starting on :%s", cfg.Port)
    if err := srv.Run(); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}

func initMetrics() {
    prometheus.MustRegister(requestDuration)
    prometheus.MustRegister(activeConnections)
}