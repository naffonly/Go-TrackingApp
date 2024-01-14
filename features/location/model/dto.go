package model

type LocationDTO struct {
	CompanyID string `json:"company_id"`
	Name      string `json:"name"`
	Lat       string `json:"lat"`
	Lon       string `json:"lon"`
	Type      string `json:"type"`
	Note      string `json:"note"`
}
