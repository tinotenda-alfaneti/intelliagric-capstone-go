package config

import (
    "fmt"
    "intelliagric-backend/internal/models"
    "log"
    "os"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

type dbEnvVars struct {
    DbUser, DbPassword, DbName, DbHost string
}

type Database struct {
    DB *gorm.DB
}

func (db *Database) getDbEnvVars() *dbEnvVars {
    return &dbEnvVars{
        DbUser:     os.Getenv("DB_USER"),
        DbPassword: os.Getenv("DB_PASSWORD"),
        DbName:     os.Getenv("DB_NAME"),
        DbHost:     os.Getenv("DB_HOST"),
    }
}

func (db *Database) Connect() {
    envVars := db.getDbEnvVars()
    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable", envVars.DbHost, envVars.DbUser, envVars.DbPassword, envVars.DbName)
    var err error
    db.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    db.DB.AutoMigrate(&models.User{})

    log.Println("Database connected successfully")
}