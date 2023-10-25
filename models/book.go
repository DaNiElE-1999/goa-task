package models

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model         // gorm.Model adds fields: ID, CreatedAt, UpdatedAt, DeletedAt
	Title       string `gorm:"uniqueIndex"` // Use tags to define constraints and mappings
	Author      string
	BookCover   string //toBeChanged
	PublishedAt time.Time
}

func MigrateBooks(db *gorm.DB) error {
	err := db.AutoMigrate(&Book{})
	return err
}
