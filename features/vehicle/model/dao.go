package model

import (
	"gorm.io/gorm"
	"time"
	company "trackingApp/features/company/model"
)

type Vehicle struct {
	ID         string          `json:"id" gorm:"primaryKey; type:varchar(255)"`
	CompanyID  string          `json:"company_id"`
	PlatNumber string          `json:"plat_number" gorm:"type:varchar(20);default:null;uniqueIndex"`
	CreateID   string          `json:"create_id"`
	CreatedAt  time.Time       `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time       `json:"updated_at" gorm:"autoCreateTime"`
	DeletedAt  gorm.DeletedAt  `gorm:"index" swaggerignore:"true" json:"deleted_at"`
	Company    company.Company `json:"company" gorm:"foreignKey:CompanyID"`
}
