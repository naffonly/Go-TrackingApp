package service

import (
	"errors"
	uuid2 "github.com/google/uuid"
	model2 "trackingApp/features/location/model"
	"trackingApp/features/order/model"
	"trackingApp/features/order/repository"
	"trackingApp/helper/pagination"
	"trackingApp/helper/random"
)

type orderServiceImpl struct {
	Repository repository.OrderInterfaceInterface
}

type OrderServiceInterface interface {
	FindAll(param pagination.QueryParam, ownerRole string, ownerId string) ([]model.Order, *pagination.Pagination, error)
	FindById(uuid string, ownerRole string, ownerId string) (*model.Order, error)
	Insert(payload *model.OrderDTO, ownerRole string, ownerId string) (*model.Order, error)
	Update(payload *model.OrderDTO, uuid string, ownerRole string, ownerId string) (*model.Order, error)
	Delete(uuid string, ownerRole string, ownerId string) error
	GetCustomerName(name string, ownerRole string, ownerId string) (*[]model.Order, error)
}

func NewOrderServiceImpl(repo repository.OrderInterfaceInterface) OrderServiceInterface {
	return &orderServiceImpl{Repository: repo}
}

func (service *orderServiceImpl) FindAll(param pagination.QueryParam, ownerRole string, ownerId string) ([]model.Order, *pagination.Pagination, error) {
	rs, err := service.Repository.FindALL(param)
	if err != nil {
		return nil, nil, errors.New("get data company failed")
	}
	var orderRes []model.Order

	for _, value := range rs {
		orderRes = append(orderRes, value)
	}

	total, err := service.Repository.TotalData()
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
	rs, err := service.Repository.FindByID(uuid)
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
		Identity:          identity,
		DropoffLocationID: uuidDrop.String(),
		PickupLocationID:  uuidPick.String(),
		CustomerName:      payload.CustomerName,
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

	rs, err := service.Repository.Insert(&newPayload)
	if err != nil {
		return nil, errors.New("failed Insert Data")
	}
	return rs, nil
}

func (service *orderServiceImpl) Update(payload *model.OrderDTO, uuid string, ownerRole string, ownerId string) (*model.Order, error) {

	newPayload := model.Order{
		CustomerName: payload.CustomerName,
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
