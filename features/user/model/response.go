package model

import (
	companyModel "trackingApp/features/company/model"
)

type UserResponse struct {
	ID        string               `json:"ID"`
	Username  string               `json:"username" `
	Name      string               `json:"name"`
	CompanyID string               `json:"company_id"`
	Email     string               `json:"email"`
	Password  string               `json:"password"`
	Role      uint                 `json:"role"`
	Company   companyModel.Company `gorm:"foreignkey:CompanyID"`
}
