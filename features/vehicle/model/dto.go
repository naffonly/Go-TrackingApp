package model

type VehicleDTO struct {
	PlatNumber string `json:"plat_number" gorm:"varchar(20);default:null"`
}
