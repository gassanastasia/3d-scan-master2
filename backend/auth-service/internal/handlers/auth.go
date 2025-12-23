package handlers

import (
	"auth-service/internal/auth"
	"auth-service/internal/tenants"
	"auth-service/pkg/models"
	"go/token"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthHandler struct {
	db        *gorm.DB
	jwtSecret string
}

func NewAuthHandler(db *gorm.DB, jwtSecret string) *AuthHandler {
	return &AuthHandler {
		db: db,
		jwtSecret: jwtSecret,
	}
}

func (h *AuthHandler) CreateTenant(c *gin.Context) {
	var req struct {
		Name string `json:"name" binding:"required"`
		AdminEmail string `json:"admin_email" binding:"required, email"`
		Password string `json:"password" binding:"required, min=6"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tenant, err := tenants.CreateTenant(h.db, req.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create tenant"})
		return
	}

	user, err := tenants.CreateUser(h.db, tenant.ID, req.AdminEmail, req.Password, "admin")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create admin user"})
		return
	}

	token, err := auth.GenerateToken(user.ID, tenant.ID, user.Email, user.Role, h.jwtSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"tenant": gin.H{
			"id": tenant.ID,
			"name": tenant.Name,
		},
		"user": gin.H{
			"id": user.ID,
			"email": user.Email,
			"role": user.Role,
		},
		"token": token,
	})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req struct {
		Email string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := h.db.Where("email = ?", req.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}


	token, err := auth.GenerateToken(user.ID, user.TenantID, user.Email, user.Role, h.jwtSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"user": gin.H{
			"id": user.ID,
			"email": user.Email,
			"role": user.Role,
		},
		"token": token,
	})
}