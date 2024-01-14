package repository

import (
	"errors"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"trackingApp/features/location/model"
	userModel "trackingApp/features/user/model"
	"trackingApp/helper/pagination"
)

type locationRepositoryImpl struct {
	DB *gorm.DB
}

func (repository *locationRepositoryImpl) TotalData() (int64, error) {
	var location model.Location
	var total int64

	result := repository.DB.Model(&location).Count(&total)
	if result.Error != nil {
		return -1, result.Error
	}

	return total, nil
}

type LocationRepositoryInterface interface {
	FindALL(param pagination.QueryParam, companyID string) ([]model.Location, error)
	FindByID(uuid string, companyID string) (*model.Location, error)
	Insert(payload *model.Location) (*model.Location, error)
	Update(payload *model.Location, uuid string) (*model.Location, error)
	Delete(uuid string) error
	TotalData() (int64, error)
	GetCompanyUser(uuid string) (*userModel.User, error)
	FIndByNote(string2 string, payload *[]model.Location) error
}

func NewLocationRepositoryImpl(db *gorm.DB) LocationRepositoryInterface {
	return &locationRepositoryImpl{DB: db}
}

func (repository *locationRepositoryImpl) GetCompanyUser(uuid string) (*userModel.User, error) {
	var user userModel.User
	err := repository.DB.Select("company_id").Where("id=?", uuid).First(&user)
	if err.Error != nil {
		return nil, err.Error
	}
	return &user, nil
}

func (repository *locationRepositoryImpl) FIndByNote(query string, payload *[]model.Location) error {
	rs := repository.DB.Where("note like ?", "%"+query+"%").Find(&payload)
	if rs.Error != nil {
		return errors.New("location note not found")
	}
	return nil
}

func (repository *locationRepositoryImpl) FindALL(param pagination.QueryParam, companyID string) ([]model.Location, error) {
	var payload []model.Location
	var offset = (param.Page - 1) * param.Size

	result := repository.DB.Preload("Company").Offset(offset).Limit(param.Size).Where("company_id=?", companyID).Find(&payload)
	if result.Error != nil {
		panic(result.Error)
		return nil, result.Error
	}
	return payload, nil
}

func (repository *locationRepositoryImpl) FindByID(uuid string, companyID string) (*model.Location, error) {
	var payload model.Location

	result := repository.DB.Preload("Company").Where("id =?", uuid).Where("company_id=?", companyID).First(&payload)
	if result.Error != nil {
		logrus.Info("Failed Find Data")
		return nil, result.Error
	}
	return &payload, nil
}

func (repository *locationRepositoryImpl) Insert(payload *model.Location) (*model.Location, error) {
	result := repository.DB.Create(&payload)
	if result.Error != nil {
		logrus.Info("Insert Data Error : ", result.Error)
		return nil, result.Error
	}
	return payload, nil
}

func (repository *locationRepositoryImpl) Update(payload *model.Location, uuid string) (*model.Location, error) {
	var location model.Location

	rs := repository.DB.Model(&location).Where("id=?", uuid).Updates(payload)
	if rs.Error != nil {
		logrus.Panic("update Data Error : ", rs.Error)
		return nil, rs.Error
	}
	return payload, nil
}

func (repository *locationRepositoryImpl) Delete(uuid string) error {
	var payload model.Location
	if err := repository.DB.Where("id=?", uuid).First(&payload).Error; err != nil {
		return errors.New("data not found")
	}

	rs := repository.DB.Delete(&payload)
	if rs.Error != nil {
		return errors.New("failed deleted data")
	}
	return nil
}
