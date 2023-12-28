package model

import (
	"gorm.io/gorm"
	"time"
	company "trackingApp/features/company/model"
	location "trackingApp/features/location/model"
)

type Order struct {
	ID                string            `json:"id" gorm:"primaryKey; type:varchar(255)"`
	CompanyID         string            `json:"company_id"`
	Identity          string            `json:"identity" gorm:"type:varchar(16);default:null"`
	CustomerName      string            `json:"customer_name" gorm:"type:varchar(100);default:null"`
	PickupLocationID  string            `json:"pickup_location_id"`
	DropoffLocationID string            `json:"dropoff_location_id"`
	CreateID          string            `json:"create_id"`
	CreatedAt         time.Time         `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt         time.Time         `json:"updated_at" gorm:"autoCreateTime"`
	DeletedAt         gorm.DeletedAt    `gorm:"index" swaggerignore:"true" json:"deleted_at"`
	Company           company.Company   `json:"company" gorm:"foreignKey:CompanyID"`
	PickupLocation    location.Location `json:"pickup_location" gorm:"foreignKey:PickupLocationID"`
	DropoffLocation   location.Location `json:"dropoff_location" gorm:"foreignKey:DropoffLocationID"`
}
