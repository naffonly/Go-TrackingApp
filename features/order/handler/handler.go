package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"trackingApp/features/order/model"
	"trackingApp/features/order/service"
	"trackingApp/helper/pagination"
	"trackingApp/utils/token"
)

type orderHandlerImpl struct {
	Service service.OrderServiceInterface
}

var (
	msg string
)

type OrderHandlerInterface interface {
	FindAll() gin.HandlerFunc
	FindByID() gin.HandlerFunc
	Insert() gin.HandlerFunc
	Update() gin.HandlerFunc
	Delete() gin.HandlerFunc
}

func NewOrderHandlerImpl(service service.OrderServiceInterface) OrderHandlerInterface {
	return &orderHandlerImpl{Service: service}
}

func (handler *orderHandlerImpl) FindAll() gin.HandlerFunc {
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
			rs, err := handler.Service.GetCustomerName(queryParam.Query, ownerRole, ownerId)
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
			c.JSON(http.StatusOK, gin.H{
				"message": "Failed",
			})
		}
		data := pagination.FormatPaginationResponse(
			"Success Find All Data",
			result,
			paginationRes,
		)
		c.JSON(http.StatusOK, data)
	}
}

func (handler *orderHandlerImpl) FindByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		uuid := c.Param("id")
		ownerId, ownerRole, err := token.ExtractTokenID(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, pagination.FormatResponse(err.Error(), nil, http.StatusBadRequest))
			c.Abort()
			return
		}
		result, err := handler.Service.FindById(uuid, ownerRole, ownerId)
		if err != nil {
			c.JSON(http.StatusBadRequest, pagination.FormatResponse(err.Error(), nil, http.StatusBadRequest))
			c.Abort()
			return
		}
		msg = "Success Find Data"
		data := pagination.FormatResponse(
			msg,
			result,
			http.StatusOK,
		)
		c.JSON(http.StatusOK, data)
	}
}

func (handler *orderHandlerImpl) Insert() gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload model.OrderDTO

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

		result, err := handler.Service.Insert(&payload, ownerRole, ownerId)
		if err != nil {
			c.JSON(http.StatusBadRequest, pagination.FormatResponse(err.Error(), nil, http.StatusBadRequest))
			c.Abort()
			return
		}
		msg = "Success Insert Data"
		data := pagination.FormatResponse(msg, result, http.StatusOK)
		c.JSON(http.StatusOK, data)

	}
}

func (handler *orderHandlerImpl) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		uuid := c.Param("id")
		var payload model.OrderDTO

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

		result, err := handler.Service.Update(&payload, uuid, ownerRole, ownerId)
		if err != nil {
			c.JSON(http.StatusBadRequest, pagination.FormatResponse(err.Error(), nil, http.StatusBadRequest))
			c.Abort()
			return
		}
		msg = "Success Update Data"
		data := pagination.FormatResponse(msg, result, http.StatusOK)
		c.JSON(http.StatusOK, data)
	}
}

func (handler *orderHandlerImpl) Delete() gin.HandlerFunc {
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
		data := pagination.FormatResponse(msg, nil, http.StatusOK)
		c.JSON(http.StatusOK, data)
	}
}
