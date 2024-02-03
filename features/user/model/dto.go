package model

type UserDto struct {
	Username  string `json:"username" validate:"required"`
	Name      string `json:"name" validate:"required"`
	Password  string `json:"password" validate:"required"`
	Email     string `json:"email" validate:"required"`
	CompanyId string `json:"company_id" validate:"required"`
	Role      uint   `json:"role" validate:"required"`
}
