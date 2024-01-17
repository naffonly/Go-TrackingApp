package mapping

import (
	"trackingApp/features/vehicle/model"
)

func VehicleToResponse(payload *model.Vehicle) *model.VehicleResponse {
	return &model.VehicleResponse{
		ID:         payload.ID,
		CompanyID:  payload.CompanyID,
		PlatNumber: payload.PlatNumber,
		CreateID:   payload.CreateID,
		Company:    payload.Company,
	}
}
func DtoToVehicle(payload *model.VehicleDTO, owner string, uuid string) *model.Vehicle {
	return &model.Vehicle{
		ID:         uuid,
		PlatNumber: payload.PlatNumber,
		CreateID:   owner,
	}
}
