package model

type VehicleDTO struct {
	PlatNumber string `json:"plat_number" validate:"required"`
}
