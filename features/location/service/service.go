package service

import (
	"trackingApp/features/location/model"
	"trackingApp/features/location/repository"
)

type locationServiceImpl struct {
	Repository repository.LocationRepositoryInterface
}

type LocationServiceInterface interface {
	FindAll(payload *model.Location) (*model.Location, error)
	FIndByID(uuid string) (*model.Location, error)
	Insert(payload *model.Location) (*model.Location, error)
	Update(payload *model.Location) (*model.Location, error)
	Delete(uuid string) error
}

func NewLocationSeriveImpl(repository repository.LocationRepositoryInterface) LocationServiceInterface {
	return &locationServiceImpl{Repository: repository}
}

func (l locationServiceImpl) FindAll(payload *model.Location) (*model.Location, error) {
	//TODO implement me
	panic("implement me")
}

func (l locationServiceImpl) FIndByID(uuid string) (*model.Location, error) {
	//TODO implement me
	panic("implement me")
}

func (l locationServiceImpl) Insert(payload *model.Location) (*model.Location, error) {
	//TODO implement me
	panic("implement me")
}

func (l locationServiceImpl) Update(payload *model.Location) (*model.Location, error) {
	//TODO implement me
	panic("implement me")
}

func (l locationServiceImpl) Delete(uuid string) error {
	//TODO implement me
	panic("implement me")
}
