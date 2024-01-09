package repository

import (
	"errors"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"log"
	userModel "trackingApp/features/user/model"
	"trackingApp/helper/pagination"
)

type userRepositoryImpl struct {
	DB *gorm.DB
}

type UserRepositoryInterface interface {
	FindAll(pagination pagination.QueryParam) ([]userModel.User, error)
	FindById(uuid string) (*userModel.User, error)
	Insert(payload *userModel.User) (*userModel.User, error)
	Update(payload *userModel.User, uuid string) (*userModel.User, error)
	Delete(uuid string) error
	GetUsername(username string, data *[]userModel.User) error
	UsernameAvailable(username string) error
	EmailAvailable(email string) error
	TotalData() (int64, error)
	GetCurrentCompany(uuid string) (*userModel.User, error)
}

func NewUserRepositoryInterface(db *gorm.DB) UserRepositoryInterface {
	return &userRepositoryImpl{DB: db}
}

func (repository *userRepositoryImpl) GetCurrentCompany(uuid string) (*userModel.User, error) {
	var model userModel.User
	rs := repository.DB.Preload("Company").Select("company_id", "name").Where("id =?", uuid).First(&model)

	if rs.Error != nil {
		logrus.Panic("Failed Find Data by id")
		return nil, rs.Error
	}
	return &model, nil
}

func (repository *userRepositoryImpl) GetUsername(username string, data *[]userModel.User) error {
	rs := repository.DB.Preload("Company").Where("username like ?", "%"+username+"%").Find(data)
	if rs.Error != nil {
		return errors.New("user not found")
	}
	return nil
}

func (repository *userRepositoryImpl) FindAll(pagination pagination.QueryParam) ([]userModel.User, error) {
	var payload []userModel.User

	var offset = (pagination.Page - 1) * pagination.Size

	result := repository.DB.Preload("Company").Offset(offset).Limit(pagination.Size).Find(&payload)
	if result.Error != nil {
		panic(result.Error)
		return nil, result.Error
	}

	return payload, nil
}

func (repository *userRepositoryImpl) UsernameAvailable(username string) error {
	var user userModel.User
	if rs := repository.DB.Where("username = ?", username).First(&user); rs.RowsAffected > 0 {
		return errors.New("username already exist")
	}
	return nil
}

func (repository *userRepositoryImpl) FindById(uuid string) (*userModel.User, error) {
	var data userModel.User
	rs := repository.DB.Preload("Company").Where("id=?", uuid).First(&data)

	if rs.Error != nil {
		log.Println("Failed Find Data by id")
		return nil, rs.Error
	}
	return &data, nil
}

func (repository *userRepositoryImpl) Insert(payload *userModel.User) (*userModel.User, error) {
	repository.DB.Create(&payload)
	return payload, nil
}

func (repository *userRepositoryImpl) Update(payload *userModel.User, uuid string) (*userModel.User, error) {
	var user userModel.User
	rs := repository.DB.Model(&user).Where("id=?", uuid).Updates(payload)
	if rs.Error != nil {
		logrus.Panic("update Data Error : ", rs.Error)
		return nil, rs.Error
	}
	return payload, nil
}

func (repository *userRepositoryImpl) Delete(uuid string) error {
	var user userModel.User
	if err := repository.DB.Where("id=?", uuid).First(&user).Error; err != nil {
		return errors.New("data not found")
	}

	rs := repository.DB.Delete(&user)
	if rs.Error != nil {
		return errors.New("failed deleted data")
	}
	return nil
}

func (repository *userRepositoryImpl) EmailAvailable(email string) error {
	var user userModel.User
	if rs := repository.DB.Where("email = ?", email).First(&user); rs.RowsAffected > 0 {
		return errors.New("email already exist")
	}
	return nil
}

func (repository *userRepositoryImpl) TotalData() (int64, error) {
	var user userModel.User
	var total int64
	result := repository.DB.Model(&user).Count(&total)
	if result.Error != nil {
		return -1, result.Error
	}

	return total, nil
}
