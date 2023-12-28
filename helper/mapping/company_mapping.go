package mapping

import (
	"gorm.io/gorm"
	"time"
	"trackingApp/features/company/model"
)

func CompanyToResponse(payload *model.Company) *model.CompanyResponse {
	return &model.CompanyResponse{
		Address: payload.Address,
		Phone:   payload.Phone,
		Name:    payload.Name,
	}
}

func DtoToCompany(payload *model.CompanyDTO, owner string, uuid string) *model.Company {
	return &model.Company{
		ID:        uuid,
		Address:   payload.Address,
		Phone:     payload.Phone,
		CreateID:  owner,
		Name:      payload.Name,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: gorm.DeletedAt{},
	}
}
