package mapping

import (
	"trackingApp/features/location/model"
)

func LocationToResponse(payload *model.Location) *model.LocationResponse {
	return &model.LocationResponse{
		ID:        payload.ID,
		CompanyID: payload.CompanyID,
		Name:      payload.Name,
		Lat:       payload.Lat,
		Lon:       payload.Lon,
		Type:      payload.Type,
		Note:      payload.Note,
	}
}

func DtoToLocation(payload *model.LocationDTO, uuid string) *model.Location {
	return &model.Location{
		ID:        uuid,
		CompanyID: payload.CompanyID,
		Name:      payload.Name,
		Lat:       payload.Lat,
		Lon:       payload.Lon,
		Type:      payload.Type,
		Note:      payload.Note,
	}
}
