package models

import "time"

type Categories struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	Name       string    `json:"name" binding:"required,min=3"`
	CreatedBy  string    `json:"created_by" binding:"required"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
	ModifiedAt time.Time `gorm:"autoUpdateTime" json:"modified_at"`
	ModifiedBy string    `json:"modified_by" binding:"required"`
}
