package models

import (
	"gorm.io/gorm"
)

type Todo struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

func MigrateTodo(DB *gorm.DB) error {
	err := DB.AutoMigrate(&Todo{})
	return err
}
