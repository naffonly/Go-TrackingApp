package repository

import (
	"gorm.io/gorm"
	"trackingApp/features/location/model"
	"trackingApp/helper/pagination"
)

type locationRepositoryImpl struct {
	Db *gorm.DB
}

type LocationRepositoryInterface interface {
	FindALL(param pagination.QueryParam) (*[]model.Location, error)
	FindByID(uuid string) (*model.Location, error)
	Insert(payload *model.Location) (*model.Location, error)
	Update(payload *model.Location) (*model.Location, error)
	Delete(uuid string) error
}

func NewLocationRepositoryImpl(db *gorm.DB) LocationRepositoryInterface {
	return &locationRepositoryImpl{Db: db}
}

func (l locationRepositoryImpl) FindALL(param pagination.QueryParam) (*[]model.Location, error) {
	//TODO implement me
	panic("implement me")
}

func (l locationRepositoryImpl) FindByID(uuid string) (*model.Location, error) {
	//TODO implement me
	panic("implement me")
}

func (l locationRepositoryImpl) Insert(payload *model.Location) (*model.Location, error) {
	//TODO implement me
	panic("implement me")
}

func (l locationRepositoryImpl) Update(payload *model.Location) (*model.Location, error) {
	//TODO implement me
	panic("implement me")
}

func (l locationRepositoryImpl) Delete(uuid string) error {
	//TODO implement me
	panic("implement me")
}
