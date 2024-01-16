package model

import (
	company "trackingApp/features/company/model"
)

type VehicleResponse struct {
	ID         string          `json:"id" gorm:"primaryKey; type:varchar(255)"`
	CompanyID  string          `json:"company_id"`
	PlatNumber string          `json:"plat_number" gorm:"varchar(20);default:null"`
	CreateID   string          `json:"create_id"`
	Company    company.Company `json:"company" gorm:"foreignKey:CompanyID"`
}
