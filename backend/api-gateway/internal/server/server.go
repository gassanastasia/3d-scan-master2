package server

import (
	"api-gateway/internal/config"
	"api-gateway/internal/handlers"

	"github.com/gin-gonic/gin"
)

type Server struct {
	config *config.Config
	router *gin.Engine
}

func NewServer(cfg *config.Config) *Server {
	if cfg.Enviroment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	return &Server{
		config: cfg,
		router: gin.Default(),
	}
}

func (s *Server) setupRoutes() {
	s.router.GET("/health", handlers.HealthCheck)

	api := s.router.Group("/api/v1")
	{
		api.GET("/status", handlers.GetStatus)
		api.POST("/scans", handlers.CreateScan)
	}
}

func (s *Server) Run() error {
	s.setupRoutes()
	return s.router.Run(":" + s.config.Port)
}
