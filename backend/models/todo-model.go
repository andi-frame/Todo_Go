package models

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Title     string    `json:"title"`
	Desc      string    `json:"desc"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}

func MigrateTodo(DB *gorm.DB) error {
	err := DB.AutoMigrate(&Todo{})
	return err
}
