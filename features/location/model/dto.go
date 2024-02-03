package model

type LocationDTO struct {
	CompanyID string `json:"company_id" validate:"required"`
	Name      string `json:"name" validate:"required"`
	Lat       string `json:"lat" validate:"required"`
	Lon       string `json:"lon" validate:"required"`
	Type      string `json:"type" validate:"required"`
	Note      string `json:"note" validate:"required"`
}
