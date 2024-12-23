package models

import "time"

type Book struct {
	ID          uint       `gorm:"primaryKey" json:"id"`
	Title       string     `json:"title" binding:"required"`
	CategoryID  uint       `json:"category_id"`
	Category    Categories `gorm:"foreignKey:CategoryID" json:"category"`
	Description string     `json:"description" binding:"required"`
	ImageURL    string     `json:"image_url"`
	ReleaseYear int        `json:"release_year"`
	Price       float64    `json:"price"`
	TotalPage   int        `json:"total_page"`
	Thickness   float64    `json:"thickness"`
	CreatedBy   string     `json:"created_by"`
	CreatedAt   time.Time  `json:"created_at"`
	ModifiedBy  string     `json:"modified_by"`
	ModifiedAt  time.Time  `json:"modified_at"`
}
