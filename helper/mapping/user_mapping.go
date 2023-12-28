package mapping

import (
	authModel "trackingApp/features/auth/model"
	"trackingApp/features/company/model"
	userModel "trackingApp/features/user/model"
)

func UserToLogin(data *userModel.User, token string) *authModel.LoginResponse {
	return &authModel.LoginResponse{
		Token:    token,
		Username: data.Username,
		Name:     data.Name,
		Email:    data.Email,
	}
}

func DtoToUser(data *userModel.UserDto, uuid string, owner string) *userModel.User {
	return &userModel.User{
		ID:        uuid,
		Username:  data.Username,
		Name:      data.Name,
		CompanyID: owner,
		Email:     data.Email,
		Password:  data.Password,
		Role:      data.Role,
	}
}
func UserToResponse(data *userModel.User) *userModel.UserResponse {
	return &userModel.UserResponse{
		Username:  data.Username,
		Name:      data.Name,
		CompanyID: data.CompanyID,
		Email:     data.Email,
		Role:      data.Role,
		Company:   model.Company{},
	}
}

func ManyUserToResponse(data *[]userModel.User) *[]userModel.UserResponse {
	return &[]userModel.UserResponse{}
}
