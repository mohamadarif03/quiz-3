package models

type Book struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Title       string `json:"title" binding:"required,min=3"`
	Author      string `json:"author" binding:"required"`
	Description string `json:"description" binding:"required"`
}
