package repository

import (
	"errors"
	"gorm.io/gorm"
	"log"
	model2 "trackingApp/features/location/model"
	"trackingApp/features/order/model"
	userModel "trackingApp/features/user/model"
	"trackingApp/helper/pagination"
)

type orderRepositoryImpl struct {
	DB *gorm.DB
}

type OrderInterfaceInterface interface {
	FindALL(param pagination.QueryParam, company string) ([]model.Order, error)
	FindByID(uuid string, company string) (*model.Order, error)
	Insert(payload *model.Order) (*model.Order, error)
	Update(payload *model.Order, uuid string) (*model.Order, error)
	Delete(uuid string) error
	GetCustomerName(name string, data *[]model.Order) error
	TotalData(company string) (int64, error)
	GetCurrentCompany(uuid string) (*userModel.User, error)
	GetIdentity(identity string, company string) (*model.Order, error)
}

func NewOrderRepositoryImpl(Db *gorm.DB) OrderInterfaceInterface {
	return &orderRepositoryImpl{DB: Db}
}

func (repo *orderRepositoryImpl) GetIdentity(identity string, company string) (*model.Order, error) {
	var payload model.Order
	rs := repo.DB.Preload("Company").Preload("Vehicle").Preload("PickupLocation").Preload("DropoffLocation").Where("identity = ?", identity).Where("company_id=?", company).First(&payload)

	if rs.Error != nil {
		return nil, errors.New("data not found")
	}
	return &payload, nil
}

func (repo *orderRepositoryImpl) GetCurrentCompany(uuid string) (*userModel.User, error) {
	var user userModel.User
	err := repo.DB.Where("id =?", uuid).First(&user)
	if err.Error != nil {
		return nil, err.Error
	}
	return &user, nil
}

func (repo *orderRepositoryImpl) TotalData(company string) (int64, error) {
	var user model.Order
	var total int64
	result := repo.DB.Model(&user).Where("company_id=?", company).Count(&total)
	if result.Error != nil {
		return -1, result.Error
	}

	return total, nil
}

func (repo *orderRepositoryImpl) GetCustomerName(name string, data *[]model.Order) error {
	rs := repo.DB.Preload("Company").Preload("Vehicle").Preload("PickupLocation").Preload("DropoffLocation").Where("customer_name like ?", "%"+name+"%").Find(&data)
	if rs.Error != nil {
		return errors.New("customer not found")
	}
	return nil
}

func (repo *orderRepositoryImpl) FindALL(param pagination.QueryParam, company string) ([]model.Order, error) {
	var payload []model.Order

	var offset = (param.Page - 1) * param.Size

	result := repo.DB.Distinct().Preload("Company").Preload("Vehicle").Preload("PickupLocation").Preload("DropoffLocation").Where("company_id=?", company).Offset(offset).Limit(param.Size).Find(&payload)
	if result.Error != nil {
		panic(result.Error)
		return nil, result.Error
	}

	return payload, nil
}

func (repo *orderRepositoryImpl) FindByID(uuid string, company string) (*model.Order, error) {
	var payload model.Order
	rs := repo.DB.Preload("Company").Preload("Vehicle").Preload("PickupLocation").Preload("DropoffLocation").Where("id = ?", uuid).Where("company_id=?", company).First(&payload)

	if rs.Error != nil {
		log.Println("Failed Find Data by id")
		return nil, rs.Error
	}
	return &payload, nil
}

func (repo *orderRepositoryImpl) Insert(payload *model.Order) (*model.Order, error) {
	repo.DB.Create(payload)
	return payload, nil
}

func (repo *orderRepositoryImpl) Update(payload *model.Order, uuid string) (*model.Order, error) {
	var order model.Order
	var pickUp model2.Location
	var dropOff model2.Location

	result := repo.DB.Where("id = ?", uuid).First(&order)
	if result.Error != nil {
		log.Println("Failed Find Data by id")
		return nil, result.Error
	}

	// Lakukan pembaruan data
	rs := repo.DB.Model(&order).Updates(&payload)
	if rs.Error != nil {
		log.Println("Failed Find Data by id")
		return nil, rs.Error
	}

	err := repo.DB.Model(&pickUp).Where("id=?", order.PickupLocationID).Updates(&payload.PickupLocation)
	if err.Error != nil {
		log.Println("Failed Find Data by id")
		return nil, err.Error
	}

	errs := repo.DB.Model(&dropOff).Where("id=?", order.DropoffLocationID).Updates(&payload.DropoffLocation)
	if errs.Error != nil {
		log.Println("Failed Find Data by id")
		return nil, errs.Error
	}

	return &order, nil
}

func (repo *orderRepositoryImpl) Delete(uuid string) error {
	var order model.Order

	rs := repo.DB.Where("id = ?", uuid).Delete(&order)
	if rs.Error != nil {
		log.Println("Failed Find Data by id")
		return rs.Error
	}
	return nil
}
