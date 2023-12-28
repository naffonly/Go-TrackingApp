package service

import (
	"golang.org/x/crypto/bcrypt"
	"trackingApp/features/auth/model"
	"trackingApp/features/auth/repository"
	userModel "trackingApp/features/user/model"
	"trackingApp/helper/mapping"
	"trackingApp/utils/password"
	"trackingApp/utils/token"
)

type AuthServiceImpl struct {
	Respository repository.AuthRepositoryInterface
}

type AuthServiceInterface interface {
	Login(payload *model.LoginDto) (*model.LoginResponse, error)
	CurrentUser(uuid string) (*userModel.User, error)
}

func NewAuthService(repo repository.AuthRepositoryInterface) AuthServiceInterface {
	return &AuthServiceImpl{Respository: repo}
}

func (service *AuthServiceImpl) Login(payload *model.LoginDto) (*model.LoginResponse, error) {
	rs, err := service.Respository.LoginCheck(payload)
	if err != nil {
		return nil, err
	}

	errs := password.VerifyPassword(payload.Password, rs.Password)
	if errs != nil && errs == bcrypt.ErrMismatchedHashAndPassword {
		return nil, err
	}

	generateToken, err := token.GenerateToken(rs.ID, rs.Role)
	if err != nil {
		return nil, err
	}

	result := mapping.UserToLogin(rs, generateToken)
	return result, nil
}

func (service *AuthServiceImpl) CurrentUser(uuid string) (*userModel.User, error) {
	rs, err := service.Respository.FindById(uuid)
	if err != nil {
		return nil, err
	}
	return rs, nil
}
