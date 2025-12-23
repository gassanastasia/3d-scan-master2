package models

import (
	"time"
	"gorm.io/gorm"
)

type Tenant struct {
	ID string `gorm:"primaryKey" json:"id"`
	Name string `gorm:"not null" json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type User struct {
	ID string `gorm:"primaryKey" json:"id"`
	Email string `gorm:"uniqueIndex;not null" json:"email"`
	Password string `gorm:"uniqueIndex;not null" json:"-"`
	TenantID string `gorm:"not null;index" json:"tenant_id"`
	Role string `gorm:"default:user" json:"role"` //admin, user?
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Tenant Tenant `gorm:"foreignKey:TenantID" json:"-"`
}