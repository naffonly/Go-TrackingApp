package model

type UserDto struct {
	Username  string `json:"username"`
	Name      string `json:"name"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	CompanyId string `json:"company_id"`
	Role      uint   `json:"role"`
}
