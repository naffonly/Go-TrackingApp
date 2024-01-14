package service

import (
	"errors"
	uuid2 "github.com/google/uuid"
	"strconv"
	"trackingApp/features/location/model"
	"trackingApp/features/location/repository"
	"trackingApp/helper/mapping"
	response "trackingApp/helper/pagination"
)

type locationServiceImpl struct {
	Repository repository.LocationRepositoryInterface
}

type LocationServiceInterface interface {
	FindAll(pagination response.QueryParam, ownerRole string, ownerId string) (*[]model.Location, *response.Pagination, error)
	FIndByID(uuid string, ownerRole string, ownerId string) (*model.Location, error)
	Insert(payload *model.LocationDTO, ownerRole string, ownerId string) (*model.Location, error)
	Update(payload *model.LocationDTO, uuid string, ownerRole string, ownerId string) (*model.LocationResponse, error)
	Delete(uuid string, ownerRole string, ownerId string) error
	FindByNote(note string, ownerRole string, ownerId string) (*[]model.Location, error)
}

func NewLocationSeriveImpl(repository repository.LocationRepositoryInterface) LocationServiceInterface {
	return &locationServiceImpl{Repository: repository}
}

func (service *locationServiceImpl) FindByNote(note string, ownerRole string, ownerId string) (*[]model.Location, error) {
	owner, _ := strconv.Atoi(ownerRole)

	company := []model.Location{}
	if owner == 3 {
		return nil, errors.New("your not allowed")
	}

	err := service.Repository.FIndByNote(note, &company)
	if err != nil {
		return nil, err
	}
	return &company, nil
}

func (service *locationServiceImpl) FindAll(pagination response.QueryParam, ownerRole string, ownerId string) (*[]model.Location, *response.Pagination, error) {
	//owner, _ := strconv.Atoi(ownerRole)
	//if owner != 1 {
	//	return nil, nil, errors.New("your not allowed")
	//}

	user, err := service.Repository.GetCompanyUser(ownerId)
	if err != nil {
		return nil, nil, err
	}

	rs, err := service.Repository.FindALL(pagination, user.CompanyID)
	if err != nil {
		return nil, nil, errors.New("get data company failed")

	}
	var locRes []model.Location

	for _, value := range rs {
		locRes = append(locRes, value)
	}

	total, err := service.Repository.TotalData()
	if err != nil {
		return nil, nil, errors.New("get total menu failed")
	}

	var DataResponse = &response.Pagination{
		Page:       pagination.Page,
		PageSize:   pagination.Size,
		TotalItems: total,
	}

	return &locRes, DataResponse, nil
}

func (service *locationServiceImpl) FIndByID(uuid string, ownerRole string, ownerId string) (*model.Location, error) {

	user, err := service.Repository.GetCompanyUser(ownerId)
	if err != nil {
		return nil, err
	}

	rs, errs := service.Repository.FindByID(uuid, user.CompanyID)
	if errs != nil {
		return nil, errs
	}

	return rs, nil
}

func (service *locationServiceImpl) Insert(payload *model.LocationDTO, ownerRole string, ownerId string) (*model.Location, error) {

	uuid, _ := uuid2.NewRandom()
	data := mapping.DtoToLocation(payload, uuid.String())

	result, err := service.Repository.Insert(data)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (service *locationServiceImpl) Update(payload *model.LocationDTO, uuid string, ownerRole string, ownerId string) (*model.LocationResponse, error) {

	newPayload := model.Location{
		Name: payload.Name,
		Lat:  payload.Lat,
		Lon:  payload.Lon,
		Type: payload.Type,
		Note: payload.Note,
	}

	rs, err := service.Repository.Update(&newPayload, uuid)
	if err != nil {
		return nil, err
	}
	result := mapping.LocationToResponse(rs)
	return result, nil
}

func (service *locationServiceImpl) Delete(uuid string, ownerRole string, ownerId string) error {
	err := service.Repository.Delete(uuid)
	if err != nil {
		return err
	}
	return nil
}
