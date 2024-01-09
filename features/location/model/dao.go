package model

import (
	"gorm.io/gorm"
	"time"
	company "trackingApp/features/company/model"
)

type Location struct {
	ID        string         `json:"id" gorm:"primaryKey; type:varchar(255)"`
	CompanyID string         `json:"company_id"`
	Lat       string         `json:"lat" gorm:"default:null"`
	Lon       string         `json:"lon" gorm:"default:null"`
	Type      string         `json:"type" gorm:"type:varchar(100);default:null"`
	Note      string         `json:"note" gorm:"type:text;default:null"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoCreateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index" swaggerignore:"true" json:"deleted_at"`

	Company company.Company `json:"company" gorm:"foreignKey:CompanyID"`
}
