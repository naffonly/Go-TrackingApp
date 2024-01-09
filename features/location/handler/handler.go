package handler

import (
	"fmt"
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
	return func(c *gin.Context) {
		fmt.Sprintf("AAA")
	}
}

func (l locationHandlerImpl) FindByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Sprintf("AAA")
	}
}

func (l locationHandlerImpl) Insert() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Sprintf("AAA")
	}
}

func (l locationHandlerImpl) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Sprintf("AAA")
	}
}

func (l locationHandlerImpl) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Sprintf("AAA")
	}
}
