package model

type VehicleDTO struct {
	CompanyID  string `json:"company_id"`
	PlatNumber string `json:"plat_number" gorm:"varchar(20);default:null"`
}
