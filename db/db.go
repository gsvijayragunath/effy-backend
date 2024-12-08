package db

import (
	"effy/gravatar-profile-card/models"
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
	AuthKey string
)

func InitDB() error {
    log.Println("Loading environment variables...")
    err := godotenv.Load("prod.env")
    if err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }
	AuthKey = os.Getenv("AUTH_KEY")
    log.Println("Connecting to database...")
    dsn := fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%s sslmode=%s",
        os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_NAME"),
        os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_SSLMODE"))

    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }
    log.Println("Database connection successful.")

    log.Println("Starting migrations...")
    if err := DB.AutoMigrate(&models.User{}, &models.Profile{}); err != nil {
        log.Fatalf("Error during migrations: %v", err)
    }
    log.Println("Database migrated successfully.")
    return nil
}
