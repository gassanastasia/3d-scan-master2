package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthCheck(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"status": "healthy",
		"service": "api-gateway",
	})
}

func GetStatus(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"version": "1.0.0",
		"message": "3d scan platform API Gateway",
	})
}

func CreateScan(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"id": "scan-001",
		"status": "created",
		"message": "Scan endpoint ready - scan sevice integration pending",
	})
}