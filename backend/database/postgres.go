package database

import (
	"fmt"
	"log"
	"os"

	"github.com/andi-frame/Todo_Go/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database", err)
	}
	fmt.Println("Database connected!")

	err = models.MigrateTodo(DB)
	if err != nil {
		log.Fatal("Failed to auto migrate database", err)
	}
	fmt.Println("Database migrated!")
}
