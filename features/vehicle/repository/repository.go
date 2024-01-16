package repository

import (
	"errors"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"log"
	userModel "trackingApp/features/user/model"
	"trackingApp/features/vehicle/model"
	"trackingApp/helper/pagination"
)

type vehicleRepositoryImpl struct {
	DB *gorm.DB
}

func (repo *vehicleRepositoryImpl) GetCompanyID(uuid string) (*userModel.User, error) {
	var user userModel.User
	rs := repo.DB.Where("id=?", uuid).First(&user)
	if rs.Error != nil {
		logrus.Panic("Errors : ", rs.Error)
		return nil, errors.New("failed find data :")
	}
	return &user, nil
}

type VehicleRpositoryInterface interface {
	FindAll(params pagination.QueryParam) ([]model.Vehicle, error)
	FindByID(uuid string) (*model.Vehicle, error)
	Insert(payload *model.Vehicle) (*model.Vehicle, error)
	Update(payload *model.Vehicle, uuid string) (*model.Vehicle, error)
	Delete(uuid string) error
	TotalData() (int64, error)
	GetPlatNumber(plat string, payload *[]model.Vehicle) error
	GetCompanyID(uuid string) (*userModel.User, error)
	ValidationPlatNumber(plat string) error
}

func NewVehicleRepositoryImpl(db *gorm.DB) VehicleRpositoryInterface {
	return &vehicleRepositoryImpl{DB: db}
}

func (repo *vehicleRepositoryImpl) ValidationPlatNumber(plat string) error {
	var data model.Vehicle
	if rs := repo.DB.Where("plat_number=?", plat).First(&data); rs.RowsAffected > 0 {
		return errors.New("plat already exist")
	}
	return nil
}

func (repo *vehicleRepositoryImpl) FindAll(params pagination.QueryParam) ([]model.Vehicle, error) {
	var vehicle []model.Vehicle
	var offset = (params.Page - 1) * params.Size

	rs := repo.DB.Preload("Company").Offset(offset).Limit(params.Size).Find(&vehicle)
	if rs.Error != nil {
		logrus.Panic("Failed get data")
		return nil, rs.Error
	}
	return vehicle, nil
}

func (repo *vehicleRepositoryImpl) FindByID(uuid string) (*model.Vehicle, error) {
	var vehicle model.Vehicle
	rs := repo.DB.Preload("Company").Where("id =?", uuid).First(&vehicle)
	if rs.Error != nil {
		log.Printf("Failed Find Data by id : %s", rs.Error)
		return nil, rs.Error
	}
	return &vehicle, nil
}

func (repo *vehicleRepositoryImpl) Insert(payload *model.Vehicle) (*model.Vehicle, error) {
	rs := repo.DB.Create(&payload)
	if rs.Error != nil {
		return nil, rs.Error
	}
	return payload, nil
}

func (repo *vehicleRepositoryImpl) Update(payload *model.Vehicle, uuid string) (*model.Vehicle, error) {
	var vehicle model.Vehicle
	rs := repo.DB.Model(&vehicle).Where("id =?", uuid).Updates(payload)
	if rs.Error != nil {
		return nil, errors.New("failed update data")
	}
	return payload, nil
}

func (repo *vehicleRepositoryImpl) Delete(uuid string) error {
	var vehicle model.Vehicle

	rs := repo.DB.Where("id=?", uuid).Delete(&vehicle)
	if rs.Error != nil {
		logrus.Panic("failed delete data")
		return rs.Error
	}
	return nil
}

func (repo *vehicleRepositoryImpl) TotalData() (int64, error) {
	var vehicle model.Vehicle
	var total int64
	rs := repo.DB.Model(&vehicle).Count(&total)

	if rs.Error != nil {
		return -1, rs.Error
	}
	return total, nil
}

func (repo *vehicleRepositoryImpl) GetPlatNumber(plat string, payload *[]model.Vehicle) error {
	rs := repo.DB.Preload("Company").Where("plat_number like ?", "%"+plat+"%").First(&payload)
	if rs.Error != nil {
		return errors.New("vehicle not found")
	}
	return nil
}
