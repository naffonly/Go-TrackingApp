package service

import (
	"errors"
	"github.com/go-playground/validator/v10"
	uuid2 "github.com/google/uuid"
	model2 "trackingApp/features/location/model"
	"trackingApp/features/order/model"
	"trackingApp/features/order/repository"
	"trackingApp/helper/pagination"
	"trackingApp/helper/random"
)

type orderServiceImpl struct {
	Repository repository.OrderInterfaceInterface
	Validation *validator.Validate
}

type OrderServiceInterface interface {
	FindAll(param pagination.QueryParam, ownerRole string, ownerId string) ([]model.Order, *pagination.Pagination, error)
	FindById(uuid string, ownerRole string, ownerId string) (*model.Order, error)
	Insert(payload *model.OrderDTO, ownerRole string, ownerId string) (*model.Order, error)
	Update(payload *model.OrderDTO, uuid string, ownerRole string, ownerId string) (*model.Order, error)
	Delete(uuid string, ownerRole string, ownerId string) error
	GetCustomerName(name string, ownerRole string, ownerId string) (*[]model.Order, error)
}

func NewOrderServiceImpl(repo repository.OrderInterfaceInterface, valid *validator.Validate) OrderServiceInterface {
	return &orderServiceImpl{
		Repository: repo,
		Validation: valid,
	}
}

func (service *orderServiceImpl) FindAll(param pagination.QueryParam, ownerRole string, ownerId string) ([]model.Order, *pagination.Pagination, error) {

	user, err := service.Repository.GetCurrentCompany(ownerId)
	if err != nil {
		return nil, nil, err
	}

	rs, errs := service.Repository.FindALL(param, user.CompanyID)
	if errs != nil {
		return nil, nil, errors.New("get data company failed")
	}
	var orderRes []model.Order

	for _, value := range rs {
		orderRes = append(orderRes, value)
	}

	total, err := service.Repository.TotalData(user.CompanyID)
	if err != nil {
		return nil, nil, errors.New("get total menu failed")
	}

	var DataResponse = &pagination.Pagination{
		Page:       param.Page,
		PageSize:   param.Size,
		TotalItems: total,
	}

	return orderRes, DataResponse, nil
}

func (service *orderServiceImpl) GetCustomerName(name string, ownerRole string, ownerId string) (*[]model.Order, error) {
	data := []model.Order{}

	err := service.Repository.GetCustomerName(name, &data)
	if err != nil {
		return nil, err
	}
	return &data, err
}

func (service *orderServiceImpl) FindById(uuid string, ownerRole string, ownerId string) (*model.Order, error) {
	user, err := service.Repository.GetCurrentCompany(ownerId)
	if err != nil {
		return nil, err
	}

	rs, err := service.Repository.FindByID(uuid, user.CompanyID)
	if err != nil {
		return nil, err
	}
	return rs, nil
}

func (service *orderServiceImpl) Insert(payload *model.OrderDTO, ownerRole string, ownerId string) (*model.Order, error) {

	identity := random.GetRandomStc()
	uuid, _ := uuid2.NewRandom()
	uuidDrop, _ := uuid2.NewRandom()
	uuidPick, _ := uuid2.NewRandom()

	newPayload := model.Order{
		ID:                uuid.String(),
		CompanyID:         payload.CompanyID,
		VehicleID:         payload.VehicleID,
		Identity:          identity,
		DropoffLocationID: uuidDrop.String(),
		PickupLocationID:  uuidPick.String(),
		Recipients:        payload.Recipients,
		CreateID:          ownerId,
		DropoffLocation: model2.Location{
			ID:        uuidDrop.String(),
			CompanyID: payload.CompanyID,
			Lat:       payload.DropoffLocation.Lat,
			Lon:       payload.DropoffLocation.Lon,
			Type:      payload.DropoffLocation.Type,
			Note:      payload.DropoffLocation.Note,
		},
		PickupLocation: model2.Location{
			ID:        uuidPick.String(),
			CompanyID: payload.CompanyID,
			Lat:       payload.PickupLocation.Lat,
			Lon:       payload.PickupLocation.Lon,
			Type:      payload.PickupLocation.Type,
			Note:      payload.PickupLocation.Note,
		},
	}
	err := service.Validation.Struct(payload)
	if err != nil {
		return nil, errors.New("validation failed please check your input and try again")
	}

	rs, err := service.Repository.Insert(&newPayload)
	if err != nil {
		return nil, errors.New("failed Insert Data")
	}
	return rs, nil
}

func (service *orderServiceImpl) Update(payload *model.OrderDTO, uuid string, ownerRole string, ownerId string) (*model.Order, error) {

	newPayload := model.Order{
		Recipients: payload.Recipients,
		VehicleID:  payload.VehicleID,
		DropoffLocation: model2.Location{
			Lat:  payload.DropoffLocation.Lat,
			Lon:  payload.DropoffLocation.Lon,
			Type: payload.DropoffLocation.Type,
			Note: payload.DropoffLocation.Note,
		},
		PickupLocation: model2.Location{
			Lat:  payload.PickupLocation.Lat,
			Lon:  payload.PickupLocation.Lon,
			Type: payload.PickupLocation.Type,
			Note: payload.PickupLocation.Note,
		},
	}
	rs, err := service.Repository.Update(&newPayload, uuid)
	if err != nil {
		return nil, errors.New("failed Update Data")
	}
	return rs, nil
}

func (service *orderServiceImpl) Delete(uuid string, ownerRole string, ownerId string) error {
	err := service.Repository.Delete(uuid)
	if err != nil {
		return err
	}
	return nil
}
