package models

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model         // gorm.Model adds fields: ID, CreatedAt, UpdatedAt, DeletedAt
	Title       string `gorm:"uniqueIndex"` // Use tags to define constraints and mappings
	Author      string
	BookCover   string //toBeChanged to support img
	PublishedAt string //toBeChanged to Date Type
}
