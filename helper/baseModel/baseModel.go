package baseModel

import (
	"gorm.io/gorm"
	"time"
)

type Base struct {
	ID        string         `gorm:"primaryKey; type:varchar(255)"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoCreateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index" swaggerignore:"true" json:"deleted_at"`
}
