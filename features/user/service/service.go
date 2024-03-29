package service

import (
	"errors"
	"github.com/go-playground/validator/v10"
	uuid2 "github.com/google/uuid"
	"strconv"
	userModel "trackingApp/features/user/model"
	"trackingApp/features/user/repository"
	"trackingApp/helper/mapping"
	response "trackingApp/helper/pagination"
	"trackingApp/utils/password"
)

type UserServiceImpl struct {
	Repository repository.UserRepositoryInterface
	Validation *validator.Validate
}

func (service *UserServiceImpl) GetCurrentCompany(ownerRole string, ownerId string) (*userModel.User, error) {
	rs, err := service.Repository.GetCurrentCompany(ownerId)
	if err != nil {
		return nil, err
	}
	return rs, nil
}

type UserServiceInterface interface {
	FindAll(pagination response.QueryParam, ownerRole string, ownerId string) ([]userModel.User, *response.Pagination, error)
	FindById(uuid string, ownerRole string, ownerId string) (*userModel.User, error)
	Insert(payload *userModel.UserDto, ownerRole string, ownerId string) (*userModel.UserResponse, error)
	Update(payload *userModel.UserDto, uuid string, ownerRole string, ownerId string) (*userModel.UserResponse, error)
	Delete(uuid string, ownerRole string, ownerId string) error
	GetUsername(username string, ownerRole string, ownerId string) (*[]userModel.User, error)
	GetCurrentCompany(ownerRole string, ownerId string) (*userModel.User, error)
}

func NewUserServiceInterface(repo repository.UserRepositoryInterface, valid *validator.Validate) UserServiceInterface {
	return &UserServiceImpl{
		Repository: repo,
		Validation: valid,
	}
}
func (service *UserServiceImpl) GetUsername(username string, ownerRole string, ownerId string) (*[]userModel.User, error) {
	owner, _ := strconv.Atoi(ownerRole)
	data := []userModel.User{}
	if owner == 3 {
		return nil, errors.New("your not allowed")
	}

	err := service.Repository.GetUsername(username, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (service *UserServiceImpl) FindAll(pagination response.QueryParam, ownerRole string, ownerId string) ([]userModel.User, *response.Pagination, error) {
	owner, _ := strconv.Atoi(ownerRole)
	if owner == 3 {
		return nil, nil, errors.New("your not allowed")
	}

	user, err := service.Repository.GetCurrentCompany(ownerId)
	if err != nil {
		return nil, nil, err
	}

	rs, err := service.Repository.FindAll(pagination, user.CompanyID)
	if err != nil {
		return nil, nil, errors.New("get data company failed")

	}
	var userRes []userModel.User

	for _, value := range rs {
		userRes = append(userRes, value)
	}

	total, err := service.Repository.TotalData(user.CompanyID)
	if err != nil {
		return nil, nil, errors.New("get total menu failed")
	}

	var DataResponse = &response.Pagination{
		Page:       pagination.Page,
		PageSize:   pagination.Size,
		TotalItems: total,
	}

	return userRes, DataResponse, nil
}

func (service *UserServiceImpl) FindById(uuid string, ownerRole string, ownerId string) (*userModel.User, error) {
	owner, _ := strconv.Atoi(ownerRole)
	if owner == 3 {
		return nil, errors.New("your not allowed")
	}

	rs, err := service.Repository.FindById(uuid)
	if err != nil {
		return nil, err
	}
	return rs, nil
}

func (service *UserServiceImpl) Insert(payload *userModel.UserDto, ownerRole string, ownerId string) (*userModel.UserResponse, error) {
	owner, _ := strconv.Atoi(ownerRole)

	if owner == 3 {
		return nil, errors.New("your not allowed")
	}

	err := service.Repository.UsernameAvailable(payload.Username)
	if err != nil {
		if err != nil {
			return nil, err
		}
	}

	errs := service.Repository.EmailAvailable(payload.Email)
	if errs != nil {
		return nil, err
	}
	id, err := uuid2.NewRandom()

	hashedPassword, err := password.HashPassword(payload.Password)
	if err != nil {
		return nil, err
	}
	payload.Role = 3
	if owner != 2 {
		payload.Role = 2
	}

	er := service.Validation.Struct(payload)
	if er != nil {
		return nil, errors.New("validation failed please check your input and try again")
	}

	newPayload := userModel.User{
		CompanyID: payload.CompanyId,
		Role:      payload.Role,
		Name:      payload.Name,
		Email:     payload.Email,
		Username:  payload.Username,
		Password:  hashedPassword,
		ID:        id.String(),
	}

	rs, err := service.Repository.Insert(&newPayload)
	if err != nil {
		return nil, errors.New("failed Insert Data")
	}

	result := mapping.UserToResponse(rs)
	return result, nil
}

func (service *UserServiceImpl) Update(payload *userModel.UserDto, uuid string, ownerRole string, ownerId string) (*userModel.UserResponse, error) {
	owner, _ := strconv.Atoi(ownerRole)
	if owner == 3 {
		return nil, errors.New("your not allowed")
	}

	newPayload := userModel.User{
		Username: payload.Username,
		Name:     payload.Name,
		Email:    payload.Email,
	}

	rs, err := service.Repository.Update(&newPayload, uuid)
	if err != nil {
		return nil, errors.New("failed Update Data")
	}
	result := mapping.UserToResponse(rs)
	return result, nil
}

func (service *UserServiceImpl) Delete(uuid string, ownerRole string, ownerId string) error {
	owner, _ := strconv.Atoi(ownerRole)
	if owner == 3 {
		return errors.New("your not allowed")
	}

	if len(uuid) == 0 {
		return errors.New("id not found")
	}
	err := service.Repository.Delete(uuid)
	return err
}
