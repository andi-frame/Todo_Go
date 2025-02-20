package database

import (
	"fmt"
	"log"

	"github.com/andi-frame/Todo_Go/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	dsn := "host=localhost user=postgres password=postgres dbname=Todo_Go_DB port=5432 sslmode=disable TimeZone=Asia/Jakarta"
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