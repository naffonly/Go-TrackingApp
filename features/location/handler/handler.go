package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"trackingApp/features/location/model"
	"trackingApp/features/location/service"
	"trackingApp/helper/pagination"
	"trackingApp/utils/token"
)

var (
	msg string
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

func (handler *locationHandlerImpl) FindAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		ownerId, ownerRole, err := token.ExtractTokenID(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, pagination.FormatResponse(err.Error(), nil, http.StatusBadRequest))
			c.Abort()
			return
		}

		var queryParam pagination.QueryParam
		queryParam.Size, _ = strconv.Atoi(c.Query("size"))
		queryParam.Page, _ = strconv.Atoi(c.Query("page"))
		queryParam.Query = c.Query("query")

		var paginationRes *pagination.Pagination
		if queryParam.Page < 1 || queryParam.Size < 1 {
			queryParam.Page = 1
			queryParam.Size = 12
		}

		if queryParam.Query != "" {
			rs, err := handler.Service.FindByNote(queryParam.Query, ownerRole, ownerId)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"errors": err.Error(),
				})
				c.Abort()
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"msg":  "data found",
				"data": rs,
			})
			c.Abort()
			return
		}

		result, paginationRes, err := handler.Service.FindAll(queryParam, ownerRole, ownerId)
		if err != nil {
			data := pagination.FormatPaginationResponse(
				err.Error(),
				result,
				paginationRes,
			)
			c.JSON(http.StatusBadRequest, data)
		}
		data := pagination.FormatPaginationResponse(
			"Success Find All Data",
			result,
			paginationRes,
		)
		c.JSON(http.StatusOK, data)
	}
}

func (handler *locationHandlerImpl) FindByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		uuid := c.Param("id")

		ownerId, ownerRole, err := token.ExtractTokenID(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, pagination.FormatResponse(err.Error(), nil, http.StatusBadRequest))
			c.Abort()
			return
		}

		if uuid == "" {
			msg = "id required"
			c.JSON(http.StatusBadRequest, pagination.FormatResponse(msg, nil, http.StatusBadRequest))
			c.Abort()
			return
		}

		result, errs := handler.Service.FIndByID(uuid, ownerRole, ownerId)
		if errs != nil {
			c.JSON(http.StatusNotFound, pagination.FormatResponse(errs.Error(), nil, http.StatusNotFound))
			c.Abort()
			return
		}

		msg = "Data Found"
		c.JSON(http.StatusOK, pagination.FormatResponse(msg, result, http.StatusOK))
	}
}

func (handler *locationHandlerImpl) Insert() gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload model.LocationDTO

		ownerId, ownerRole, err := token.ExtractTokenID(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, pagination.FormatResponse(err.Error(), nil, http.StatusBadRequest))
			c.Abort()
			return
		}

		errs := c.Bind(&payload)
		if errs != nil {
			c.JSON(http.StatusBadRequest, pagination.FormatResponse(err.Error(), nil, http.StatusBadRequest))
			c.Abort()
			return
		}

		rs, err := handler.Service.Insert(&payload, ownerRole, ownerId)
		if err != nil {
			c.JSON(http.StatusBadRequest, pagination.FormatResponse(err.Error(), nil, http.StatusBadRequest))
			c.Abort()
			return
		}

		msg = "Success Created"
		c.JSON(http.StatusOK, pagination.FormatResponse(msg, rs, http.StatusOK))
	}
}

func (handler *locationHandlerImpl) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload model.LocationDTO
		uuid := c.Param("id")
		ownerId, ownerRole, err := token.ExtractTokenID(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, pagination.FormatResponse(err.Error(), nil, http.StatusBadRequest))
			c.Abort()
			return
		}

		if errs := c.Bind(&payload); errs != nil {
			c.JSON(http.StatusBadRequest, pagination.FormatResponse(errs.Error(), nil, http.StatusBadRequest))
			c.Abort()
			return
		}

		result, err := handler.Service.Update(&payload, uuid, ownerRole, ownerId)
		if err != nil {
			c.JSON(http.StatusBadRequest, pagination.FormatResponse(err.Error(), nil, http.StatusBadRequest))
			c.Abort()
			return
		}

		msg = "Success Updating Data"
		c.JSON(http.StatusOK, pagination.FormatResponse(msg, result, http.StatusOK))
	}
}

func (handler *locationHandlerImpl) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		uuid := c.Param("id")
		ownerId, ownerRole, err := token.ExtractTokenID(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, pagination.FormatResponse(err.Error(), nil, http.StatusBadRequest))
			c.Abort()
			return
		}

		errs := handler.Service.Delete(uuid, ownerRole, ownerId)
		if errs != nil {
			c.JSON(http.StatusBadRequest, pagination.FormatResponse(err.Error(), nil, http.StatusBadRequest))
			c.Abort()
			return
		}

		msg = "Success Delete Data"
		c.JSON(http.StatusOK, pagination.FormatResponse(msg, nil, http.StatusOK))
	}
}
