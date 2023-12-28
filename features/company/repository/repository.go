package repository

import (
	"errors"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"log"
	"trackingApp/features/company/model"
	"trackingApp/helper/mapping"
	"trackingApp/helper/pagination"
)

var (
	company model.Company
)

type companyRepositoryImpl struct {
	DB *gorm.DB
}

type CompanyRepositoryInterface interface {
	Insert(payload *model.Company) (*model.CompanyResponse, error)
	Update(payload *model.Company, uuid string) (*model.Company, error)
	Delete(uuid string) error
	CompanyAvailable(name string) error
	GetCompanyName(name string, daa *[]model.Company) error
	FindById(uuid string) (*model.Company, error)
	FindAll(pagination pagination.QueryParam) ([]model.Company, error)
	TotalData() (int64, error)
}

func NewCompanyRepositoryImpl(db *gorm.DB) CompanyRepositoryInterface {
	return &companyRepositoryImpl{DB: db}
}

func (repository *companyRepositoryImpl) GetCompanyName(name string, data *[]model.Company) error {
	rs := repository.DB.Where("name like ?", "%"+name+"%").Find(&data)
	if rs.Error != nil {
		return errors.New("company not found")
	}
	return nil
}

func (repository *companyRepositoryImpl) CompanyAvailable(name string) error {
	if rs := repository.DB.Where("name = ?", name).First(&company); rs.RowsAffected > 0 {
		return errors.New("company already exist")
	}
	return nil
}

func (repository *companyRepositoryImpl) TotalData() (int64, error) {
	var total int64
	result := repository.DB.Model(&company).Count(&total)
	if result.Error != nil {
		return -1, result.Error
	}

	return total, nil
}

func (repository *companyRepositoryImpl) Insert(payload *model.Company) (*model.CompanyResponse, error) {
	repository.DB.Create(&payload)
	rs := mapping.CompanyToResponse(payload)
	return rs, nil
}

func (repository *companyRepositoryImpl) Update(payload *model.Company, uuid string) (*model.Company, error) {
	rs := repository.DB.Model(&company).Where("id=?", uuid).Updates(payload)
	if rs.Error != nil {
		logrus.Panic("update Data Error : ", rs.Error)
		return nil, rs.Error
	}
	return payload, nil
}

func (repository *companyRepositoryImpl) Delete(uuid string) error {
	if err := repository.DB.Where("id=?", uuid).First(&company).Error; err != nil {
		return errors.New("data not found")
	}

	rs := repository.DB.Delete(&company)
	if rs.Error != nil {
		return errors.New("failed deleted data")
	}
	return nil
}

func (repository *companyRepositoryImpl) FindById(uuid string) (*model.Company, error) {
	var data *model.Company
	rs := repository.DB.Where("id=?", uuid).First(&data)

	if rs.Error != nil {
		log.Println("Failed Find Data by id")
		return nil, rs.Error
	}
	return data, nil
}

func (repository *companyRepositoryImpl) FindAll(pagination pagination.QueryParam) ([]model.Company, error) {
	var payload []model.Company

	var offset = (pagination.Page - 1) * pagination.Size

	result := repository.DB.Offset(offset).Limit(pagination.Size).Find(&payload)
	if result.Error != nil {
		panic(result.Error)
		return nil, result.Error
	}

	return payload, nil
}
