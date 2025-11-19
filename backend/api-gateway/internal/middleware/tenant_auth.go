package middleware

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func TenantAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tenantID := c.GetHeader("X-Tenant-ID")
		userID := c.GetHeader("X-User-ID")

		if tenantID == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error" : "X-Tenant-ID header required",
			})

			c.Abort()
		}

		c.Set("tenant_id", tenantID)
		c.Set("user_id", userID)

		c.Next()
	}
}

func RequireTenant() gin.HandlerFunc {
	return func(c *gin.Context) {
		tenantID, exists := c.Get("tenant_id")
		if !exists || tenantID == ""{
			c.JSON(http.StatusUnauthorized, gin.H{
				"error" : "Tenant authentication required",
			})

			c.Abort()
			return
		}
		c.Next()
	}
}