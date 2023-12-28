package service

import (
	"errors"
	uuid2 "github.com/google/uuid"
	"strconv"
	"trackingApp/features/company/model"
	"trackingApp/features/company/repository"
	"trackingApp/helper/mapping"
	response "trackingApp/helper/pagination"
)

type companyServiceImpl struct {
	repo repository.CompanyRepositoryInterface
}

type CompanyServiceInterface interface {
	Insert(payload *model.CompanyDTO, ownerRole string, ownerId string) (*model.CompanyResponse, error)
	Update(payload *model.CompanyDTO, uuid string, ownerRole string, ownerId string) (*model.CompanyResponse, error)
	Delete(uuid string, ownerRole string, ownerId string) error
	FindById(uuid string, ownerRole string, ownerId string) (*model.Company, error)
	GetCompanyName(name string, ownerRole string, ownerId string) (*[]model.Company, error)
	FindAll(pagination response.QueryParam, ownerRole string, ownerId string) ([]model.Company, *response.Pagination, error)
}

func NewCompanyServiceInterface(repository repository.CompanyRepositoryInterface) CompanyServiceInterface {
	return &companyServiceImpl{repo: repository}
}

func (c *companyServiceImpl) GetCompanyName(name string, ownerRole string, ownerId string) (*[]model.Company, error) {
	owner, _ := strconv.Atoi(ownerRole)

	company := []model.Company{}
	if owner == 3 {
		return nil, errors.New("your not allowed")
	}

	err := c.repo.GetCompanyName(name, &company)
	if err != nil {
		return nil, err
	}
	return &company, nil
}

func (c *companyServiceImpl) Insert(payload *model.CompanyDTO, ownerRole string, ownerId string) (*model.CompanyResponse, error) {
	owner, _ := strconv.Atoi(ownerRole)
	if owner != 1 {
		return nil, errors.New("your not allowed")
	}

	rs := c.repo.CompanyAvailable(payload.Name)
	if rs != nil {
		return nil, rs
	}

	uuid, _ := uuid2.NewRandom()
	data := mapping.DtoToCompany(payload, ownerId, uuid.String())

	result, err := c.repo.Insert(data)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *companyServiceImpl) Update(payload *model.CompanyDTO, uuid string, ownerRole string, ownerId string) (*model.CompanyResponse, error) {
	owner, _ := strconv.Atoi(ownerRole)
	if owner != 1 {
		return nil, errors.New("your not allowed")
	}

	newPayload := model.Company{
		Name:     payload.Name,
		Address:  payload.Address,
		Phone:    payload.Phone,
		CreateID: ownerId,
	}

	rs, err := c.repo.Update(&newPayload, uuid)
	if err != nil {
		return nil, errors.New("failed Update Data")
	}
	result := mapping.CompanyToResponse(rs)
	return result, nil
}

func (c *companyServiceImpl) Delete(uuid string, ownerRole string, ownerId string) error {
	owner, _ := strconv.Atoi(ownerRole)
	if owner != 1 {
		return errors.New("your not allowed")
	}

	if len(uuid) == 0 {
		return errors.New("id not found")
	}
	err := c.repo.Delete(uuid)
	return err
}

func (c *companyServiceImpl) FindById(uuid string, ownerRole string, ownerId string) (*model.Company, error) {
	owner, _ := strconv.Atoi(ownerRole)
	if owner != 1 {
		return nil, errors.New("your not allowed")
	}

	rs, err := c.repo.FindById(uuid)
	if err != nil {
		return nil, err
	}
	return rs, nil
}

func (c *companyServiceImpl) FindAll(pagination response.QueryParam, ownerRole string, ownerId string) ([]model.Company, *response.Pagination, error) {
	owner, _ := strconv.Atoi(ownerRole)
	if owner != 1 {
		return nil, nil, errors.New("your not allowed")
	}

	rs, err := c.repo.FindAll(pagination)
	if err != nil {
		return nil, nil, errors.New("get data company failed")

	}
	var logisticRes []model.Company

	for _, value := range rs {
		logisticRes = append(logisticRes, value)
	}

	total, err := c.repo.TotalData()
	if err != nil {
		return nil, nil, errors.New("get total menu failed")
	}

	var DataResponse = &response.Pagination{
		Page:       pagination.Page,
		PageSize:   pagination.Size,
		TotalItems: total,
	}

	return logisticRes, DataResponse, nil
}
