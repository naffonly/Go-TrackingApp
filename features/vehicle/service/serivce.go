package service

import (
	"errors"
	uuid2 "github.com/google/uuid"
	"strings"
	"trackingApp/features/vehicle/model"
	"trackingApp/features/vehicle/repository"
	"trackingApp/helper/mapping"
	"trackingApp/helper/pagination"
)

type vehicleServiceImpl struct {
	repository repository.VehicleRpositoryInterface
}

type VehicleServiceInterface interface {
	FindAll(param pagination.QueryParam, ownerRole string, ownerId string) ([]model.Vehicle, *pagination.Pagination, error)
	FindByPlatNumber(platNumber string, ownerRole string, ownerId string) (*[]model.Vehicle, error)
	FindById(uuid string, ownerRole string, ownerId string) (*model.Vehicle, error)
	Insert(payload *model.VehicleDTO, ownerRole string, ownerId string) (*model.VehicleResponse, error)
	Update(payload *model.VehicleDTO, uuid string, ownerRole string, ownerId string) (*model.Vehicle, error)
	Delete(uuid string, ownerRole string, ownerId string) error
}

func NewVehicleServiceImpl(repository repository.VehicleRpositoryInterface) VehicleServiceInterface {
	return &vehicleServiceImpl{repository: repository}
}

func (service *vehicleServiceImpl) FindAll(param pagination.QueryParam, ownerRole string, ownerId string) ([]model.Vehicle, *pagination.Pagination, error) {

	user, err := service.repository.GetCompanyUser(ownerId)
	if err != nil {
		return nil, nil, err
	}

	rs, err := service.repository.FindAll(param, user.CompanyID)
	if err != nil {
		return nil, nil, errors.New("get data failed")
	}
	var dataVehicle []model.Vehicle
	for _, value := range rs {
		dataVehicle = append(dataVehicle, value)
	}
	total, err := service.repository.TotalData(user.CompanyID)
	if err != nil {
		return nil, nil, errors.New("get total data failed")
	}
	var DataResponse = &pagination.Pagination{
		Page:       param.Page,
		PageSize:   param.Size,
		TotalItems: total,
	}
	return dataVehicle, DataResponse, nil
}

func (service *vehicleServiceImpl) FindByPlatNumber(platNumber string, ownerRole string, ownerId string) (*[]model.Vehicle, error) {
	data := []model.Vehicle{}
	err := service.repository.GetPlatNumber(platNumber, &data)
	if err != nil {
		return nil, errors.New("failed find data")
	}
	return &data, nil
}

func (service *vehicleServiceImpl) FindById(uuid string, ownerRole string, ownerId string) (*model.Vehicle, error) {
	user, err := service.repository.GetCompanyUser(ownerId)
	if err != nil {
		return nil, err
	}

	rs, err := service.repository.FindByID(uuid, user.CompanyID)
	if err != nil {
		return nil, errors.New("find data failed")
	}
	return rs, nil
}

func (service *vehicleServiceImpl) Insert(payload *model.VehicleDTO, ownerRole string, ownerId string) (*model.VehicleResponse, error) {
	uuid, _ := uuid2.NewRandom()

	user, err := service.repository.GetCompanyUser(ownerId)
	if err != nil {
		return nil, err
	}

	newPayload := model.Vehicle{
		ID:         uuid.String(),
		CompanyID:  user.CompanyID,
		PlatNumber: strings.ToUpper(payload.PlatNumber),
		CreateID:   ownerId,
	}

	errs := service.repository.ValidationPlatNumber(newPayload.PlatNumber)
	if errs != nil {
		return nil, errs
	}

	rs, err := service.repository.Insert(&newPayload)
	if err != nil {
		return nil, errors.New("failed insert data")
	}
	dataRes := mapping.VehicleToResponse(rs)

	return dataRes, nil
}

func (service *vehicleServiceImpl) Update(payload *model.VehicleDTO, uuid string, ownerRole string, ownerId string) (*model.Vehicle, error) {

	newPayload := &model.Vehicle{
		PlatNumber: strings.ToUpper(payload.PlatNumber),
	}

	errs := service.repository.ValidationPlatNumber(newPayload.PlatNumber)
	if errs != nil {
		return nil, errs
	}

	rs, err := service.repository.Update(newPayload, uuid)
	if err != nil {
		return nil, errors.New("failed update data")
	}
	return rs, nil
}

func (service *vehicleServiceImpl) Delete(uuid string, ownerRole string, ownerId string) error {
	errs := service.repository.Delete(uuid)
	if errs != nil {
		return errors.New("failed delete data")
	}
	return nil
}
