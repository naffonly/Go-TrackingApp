package model

import location "trackingApp/features/location/model"

type OrderDTO struct {
	CompanyID       string            `json:"company_id"`
	CustomerName    string            `json:"customer_name"`
	PickupLocation  location.Location `json:"pickup_location"`
	DropoffLocation location.Location `json:"dropoff_location"`
}
