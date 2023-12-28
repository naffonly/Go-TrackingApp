package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"trackingApp/features/company/model"
	"trackingApp/features/company/service"
	"trackingApp/helper/pagination"
	"trackingApp/utils/token"
)

type companyHandlerImpl struct {
	Service service.CompanyServiceInterface
}

var (
	msg string
)

type CompanyHandlerInterface interface {
	Insert() gin.HandlerFunc
	Update() gin.HandlerFunc
	Delete() gin.HandlerFunc
	FindByID() gin.HandlerFunc
	FindAll() gin.HandlerFunc
}

func NewCompanyHanlderImpl(service service.CompanyServiceInterface) CompanyHandlerInterface {
	return &companyHandlerImpl{Service: service}
}

func (handler *companyHandlerImpl) Insert() gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload model.CompanyDTO
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

func (handler *companyHandlerImpl) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		uuid := c.Param("id")
		ownerId, ownerRole, err := token.ExtractTokenID(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, pagination.FormatResponse(err.Error(), nil, http.StatusBadRequest))
			c.Abort()
			return
		}

		var payload model.CompanyDTO
		if err := c.Bind(&payload); err != nil {
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

		msg = "Success Updating Data"
		c.JSON(http.StatusOK, pagination.FormatResponse(msg, result, http.StatusOK))
	}
}

func (handler *companyHandlerImpl) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		ownerId, ownerRole, err := token.ExtractTokenID(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, pagination.FormatResponse(err.Error(), nil, http.StatusBadRequest))
			c.Abort()
			return
		}

		errs := handler.Service.Delete(id, ownerRole, ownerId)
		if errs != nil {
			c.JSON(http.StatusBadRequest, pagination.FormatResponse(err.Error(), nil, http.StatusBadRequest))
			c.Abort()
			return
		}
		msg = "Success Delete Data"
		c.JSON(http.StatusOK, pagination.FormatResponse(msg, nil, http.StatusOK))
	}
}

func (handler *companyHandlerImpl) FindByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		ownerId, ownerRole, err := token.ExtractTokenID(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, pagination.FormatResponse(err.Error(), nil, http.StatusBadRequest))
			c.Abort()
			return
		}

		if id == "" {
			msg = "id required"
			c.JSON(http.StatusBadRequest, pagination.FormatResponse(msg, nil, http.StatusBadRequest))
			c.Abort()
			return
		}
		result, err := handler.Service.FindById(id, ownerRole, ownerId)
		if err != nil {
			c.JSON(http.StatusBadRequest, pagination.FormatResponse(err.Error(), nil, http.StatusBadRequest))
			c.Abort()
			return
		}
		msg = "Data Found"
		c.JSON(http.StatusOK, pagination.FormatResponse(msg, result, http.StatusOK))
	}
}

func (handler *companyHandlerImpl) FindAll() gin.HandlerFunc {
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
			rs, err := handler.Service.GetCompanyName(queryParam.Query, ownerRole, ownerId)
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
				"message": "success",
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
