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
		"features" : []string{
			"multi-tenant",
			"clickhouse-analytics",
			"klipper-integration",
			"real-time-monitoring",
		},
	})
}

func CreateScan(c *gin.Context){
	tenantID := c.GetHeader("X-Tenant-ID")
	userID := c.GetHeader("X-User-ID")

	c.JSON(http.StatusOK, gin.H{
		"id": "scan-001",
		"tenant_id": tenantID,
		"user_id": userID,
		"status": "created",
		"message": "Scan created in multi-tenant enviroment",
	})
}

func CreateTenant(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Tenant creation endpoint",
		"status": "pending_auth_service",
	})
}

func GetPrinters(c *gin.Context) {
	tenantID := c.GetHeader("X-Tenant-ID")

	c.JSON(http.StatusOK, gin.H{
		"tenant": tenantID,
		"printers": []gin.H{
			{
				"id": "printer-1",
				"name": "Klipper-1",
				"status": "online",
			},
		},
	})
}