package models

import "time"

type User struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	Username   string    `json:"username" binding:"required,min=3"`
	Password   string    `json:"password" binding:"required,min=6"`
	CreatedBy  string    `json:"created_by" binding:"required"`
	CreatedAt  time.Time `json:"created_at" binding:"required"`
	ModifiedAt time.Time `json:"modified_at" binding:"required"`
	ModifiedBy string    `json:"modified_by" binding:"required"`
}
