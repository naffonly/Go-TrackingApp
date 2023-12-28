package model

import (
	"gorm.io/gorm"
	"time"
)

type Company struct {
	ID        string         `json:"id" gorm:"primaryKey; type:varchar(255)"`
	Address   string         `json:"address" gorm:"varchar(100);Not null"`
	Phone     string         `json:"phone" gorm:"varchar(100);Not null"`
	CreateID  string         `json:"create_id"`
	Name      string         `json:"name" gorm:"varchar(100);Not null"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoCreateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index" swaggerignore:"true" json:"deleted_at"`
}
