package model

type LoginResponse struct {
	Token    string `json:"token"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
}
