package server

import (
	"api-gateway/internal/config"
	"api-gateway/internal/handlers"
	"api-gateway/internal/middleware"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Server struct {
	config *config.Config
	router *gin.Engine
}

func NewServer(cfg *config.Config) *Server {
	if cfg.Enviroment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	server := &Server{
		config: cfg,
		router: gin.Default(),
	}

	server.setupMiddleware()
	server.setupRoutes()

	return server
}

func (s *Server) setupMiddleware() {
	s.router.Use(middleware.Logger())
	s.router.Use(middleware.CORS())
	s.router.Use(middleware.RequestID())
	s.router.Use(middleware.TenantAuth())
}

func (s *Server) setupRoutes() {
	
	s.router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	s.router.GET("/health", handlers.HealthCheck)

	api := s.router.Group("/api/v2")
	{
		api.GET("/status", handlers.GetStatus)

		tenants := api.Group("/tenants")
		{
			tenants.POST("/", handlers.CreateTenant)
		}

		scans := api.Group("/scans")
		scans.Use(middleware.RequireTenant())
		{
			scans.POST("/", handlers.CreateScan)
		}

		printers := api.Group("/printers")
		printers.Use(middleware.RequireTenant())
		{
			printers.GET("/", handlers.GetPrinters)
		}
	}
}

func (s *Server) Run() error {
	s.setupRoutes()
	return s.router.Run(":" + s.config.Port)
}
