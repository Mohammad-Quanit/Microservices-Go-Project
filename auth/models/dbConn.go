package models

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// Global instance DB connection
var DB *gorm.DB

// Initialize DB connection with PostgreSQL
func InitDB(cfg Config) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", cfg.Host, cfg.User, cfg.Password, cfg.DBName, cfg.Port, cfg.SSLMode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// Migrates the User model by calling AutoMigrate.
	// This creates the necessary table in the database if it doesn't exist.
	if err := db.AutoMigrate(&User{}); err != nil {
		panic(err)
	}

	fmt.Println("Migrated database")

	DB = db
}
