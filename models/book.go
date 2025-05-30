package models

import (
	"time"

	"gorm.io/gorm"
)

// Book represents the data structure for a book
type Book struct {
    ID        uint           `gorm:"primarykey" json:"id"`
    Title     string         `json:"title"`
    Author    string         `json:"author"`
    CreatedAt time.Time      `json:"created_at"`
    UpdatedAt time.Time      `json:"updated_at"`
    DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
