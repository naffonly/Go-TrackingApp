package handler

import (
	"github.com/gin-gonic/gin"
	"trackingApp/features/location/service"
)

type locationHandlerImpl struct {
	Service service.LocationServiceInterface
}

type LocationHandlerInterface interface {
	FindAll() gin.HandlerFunc
	FindByID() gin.HandlerFunc
	Insert() gin.HandlerFunc
	Update() gin.HandlerFunc
	Delete() gin.HandlerFunc
}

func NewLocationHandlerInterface(service service.LocationServiceInterface) LocationHandlerInterface {
	return &locationHandlerImpl{Service: service}
}

func (l locationHandlerImpl) FindAll() gin.HandlerFunc {
	//TODO implement me
	panic("implement me")
}

func (l locationHandlerImpl) FindByID() gin.HandlerFunc {
	//TODO implement me
	panic("implement me")
}

func (l locationHandlerImpl) Insert() gin.HandlerFunc {
	//TODO implement me
	panic("implement me")
}

func (l locationHandlerImpl) Update() gin.HandlerFunc {
	//TODO implement me
	panic("implement me")
}

func (l locationHandlerImpl) Delete() gin.HandlerFunc {
	//TODO implement me
	panic("implement me")
}
