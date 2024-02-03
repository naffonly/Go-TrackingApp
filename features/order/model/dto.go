package model

import location "trackingApp/features/location/model"

type OrderDTO struct {
	CompanyID       string            `json:"company_id" validate:"required"`
	Recipients      string            `json:"recipients" validate:"required"`
	VehicleID       string            `json:"vehicle_id" validate:"required"`
	PickupLocation  location.Location `json:"pickup_location" validate:"required"`
	DropoffLocation location.Location `json:"dropoff_location" validate:"required"`
}
