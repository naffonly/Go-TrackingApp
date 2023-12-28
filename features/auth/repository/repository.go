package repository

import (
	"errors"
	"gorm.io/gorm"
	"trackingApp/features/auth/model"
	userModel "trackingApp/features/user/model"
)

type AuthRepositoryImpl struct {
	DB *gorm.DB
}

type AuthRepositoryInterface interface {
	FindById(uuid string) (*userModel.User, error)
	LoginCheck(payload *model.LoginDto) (*userModel.User, error)
}

func NewAuthRepository(db *gorm.DB) AuthRepositoryInterface {
	return &AuthRepositoryImpl{DB: db}
}

func (repository *AuthRepositoryImpl) FindById(uuid string) (*userModel.User, error) {
	var user userModel.User
	if err := repository.DB.Where("id=?", uuid).First(&user).Error; err != nil {
		return &user, err
	}
	user.PrepareGive()
	return &user, nil
}

func (repository *AuthRepositoryImpl) LoginCheck(payload *model.LoginDto) (*userModel.User, error) {
	var user userModel.User
	err := repository.DB.Model(userModel.User{}).Where("username =? ", payload.Username).Take(&user).Error
	if err != nil {
		return nil, errors.New("username or password is incorrect")
	}
	return &user, nil
}
