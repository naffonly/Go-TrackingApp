package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"trackingApp/features/vehicle/model"
	"trackingApp/features/vehicle/service"
	response "trackingApp/helper/pagination"
	"trackingApp/utils/token"
)

var (
	msg string
)

type vehicleHandlerImpl struct {
	Service service.VehicleServiceInterface
}

type VehicleHandlerInterface interface {
	Insert() gin.HandlerFunc
	Update() gin.HandlerFunc
	Delete() gin.HandlerFunc
	FindByID() gin.HandlerFunc
	FindAll() gin.HandlerFunc
}

func NewVehicleHandlerImpl(service service.VehicleServiceInterface) VehicleHandlerInterface {
	return &vehicleHandlerImpl{Service: service}
}

func (handler *vehicleHandlerImpl) Insert() gin.HandlerFunc {
	return func(c *gin.Context) {
		ownerId, ownerRole, err := token.ExtractTokenID(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, response.FormatResponse(err.Error(), nil, http.StatusBadRequest))
			c.Abort()
			return
		}
		var payload model.VehicleDTO
		errs := c.Bind(&payload)
		if errs != nil {
			c.JSON(http.StatusBadRequest, response.FormatResponse(err.Error(), nil, http.StatusBadRequest))
			c.Abort()
			return
		}

		result, err := handler.Service.Insert(&payload, ownerRole, ownerId)
		if err != nil {
			c.JSON(http.StatusBadRequest, response.FormatResponse(err.Error(), nil, http.StatusBadRequest))
			c.Abort()
			return
		}
		msg = "success insert data"
		c.JSON(http.StatusOK, response.FormatResponse(msg, result, http.StatusOK))
	}
}

func (handler *vehicleHandlerImpl) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		ownerId, ownerRole, err := token.ExtractTokenID(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, response.FormatResponse(err.Error(), nil, http.StatusBadRequest))
			c.Abort()
			return
		}
		uuid := c.Param("id")
		var paylaod model.VehicleDTO
		errs := c.Bind(&paylaod)
		if errs != nil {
			c.JSON(http.StatusBadRequest, response.FormatResponse(err.Error(), nil, http.StatusBadRequest))
		}

		result, err := handler.Service.Update(&paylaod, uuid, ownerRole, ownerId)
		if err != nil {
			c.JSON(http.StatusBadRequest, response.FormatResponse(err.Error(), nil, http.StatusBadRequest))
			c.Abort()
			return
		}
		msg = "success update"
		c.JSON(http.StatusOK, response.FormatResponse(msg, result, http.StatusOK))
	}
}

func (handler *vehicleHandlerImpl) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		ownerId, ownerRole, err := token.ExtractTokenID(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, response.FormatResponse(err.Error(), nil, http.StatusBadRequest))
			c.Abort()
			return
		}
		uuid := c.Param("id")
		errs := handler.Service.Delete(uuid, ownerRole, ownerId)
		if errs != nil {
			c.JSON(http.StatusBadRequest, response.FormatResponse(err.Error(), nil, http.StatusBadRequest))
			c.Abort()
			return
		}
		msg = "Success Delete Data"
		c.JSON(http.StatusOK, response.FormatResponse(msg, nil, http.StatusOK))
	}
}

func (handler *vehicleHandlerImpl) FindByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		ownerId, ownerRole, err := token.ExtractTokenID(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, response.FormatResponse(err.Error(), nil, http.StatusBadRequest))
			c.Abort()
			return
		}
		uuid := c.Param("id")
		rs, err := handler.Service.FindById(uuid, ownerRole, ownerId)
		if err != nil {
			c.JSON(http.StatusBadRequest, response.FormatResponse(err.Error(), nil, http.StatusBadRequest))
			c.Abort()
			return
		}
		msg = "Data found"
		c.JSON(http.StatusOK, response.FormatResponse(msg, rs, http.StatusOK))
	}
}

func (handler *vehicleHandlerImpl) FindAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		ownerId, ownerRole, err := token.ExtractTokenID(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, response.FormatResponse(err.Error(), nil, http.StatusBadRequest))
			c.Abort()
			return
		}

		var queryParam response.QueryParam
		queryParam.Query = c.Query("query")
		queryParam.Size, _ = strconv.Atoi(c.Query("size"))
		queryParam.Page, _ = strconv.Atoi(c.Query("size"))

		var paginationRes *response.Pagination

		if queryParam.Page < 1 || queryParam.Size < 1 {
			queryParam.Page = 1
			queryParam.Size = 12
		}

		if queryParam.Query != "" {
			rs, errs := handler.Service.FindByPlatNumber(queryParam.Query, ownerRole, ownerId)
			if errs != nil {
				c.JSON(http.StatusBadRequest, response.FormatResponse(errs.Error(), nil, http.StatusBadRequest))
				c.Abort()
				return
			}
			msg = "Data founded"
			c.JSON(http.StatusOK, response.FormatResponse(msg, rs, http.StatusOK))
		}

		result, paginationRes, err := handler.Service.FindAll(queryParam, ownerRole, ownerId)
		if err != nil {
			c.JSON(http.StatusBadRequest, response.FormatResponse(err.Error(), nil, http.StatusBadRequest))
			c.Abort()
			return
		}

		msg = "success Find All Data"
		data := response.FormatPaginationResponse(msg, result, paginationRes)
		c.JSON(http.StatusOK, data)
	}
}
