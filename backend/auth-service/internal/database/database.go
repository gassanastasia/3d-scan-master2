package database

import (
    "log"
    "auth-service/pkg/models"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

func Connect(databaseURL string) (*gorm.DB, error) {
    db, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
    if err != nil {
        return nil, err
    }

    // Автомиграция
    err = db.AutoMigrate(&models.Tenant{}, &models.User{})
    if err != nil {
        return nil, err
    }

    log.Println("Database connected and migrated")
    return db, nil
}