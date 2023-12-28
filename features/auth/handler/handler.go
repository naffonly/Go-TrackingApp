package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"trackingApp/features/auth/model"
	"trackingApp/features/auth/service"
	"trackingApp/helper/pagination"
	"trackingApp/utils/token"
)

var (
	msg string
)

type authHadlerImpl struct {
	Service service.AuthServiceInterface
}

type AuthHandlerInterface interface {
	Login() gin.HandlerFunc
	CurrentUser() gin.HandlerFunc
}

func NewAuthHandler(service service.AuthServiceInterface) AuthHandlerInterface {
	return &authHadlerImpl{Service: service}
}

func (a *authHadlerImpl) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload model.LoginDto
		err := c.Bind(&payload)
		if err != nil {
			c.JSON(http.StatusBadRequest, pagination.FormatResponse(err.Error(), nil, http.StatusBadRequest))
			c.Abort()
			return
		}

		rs, err := a.Service.Login(&payload)
		if err != nil {
			c.JSON(http.StatusBadRequest, pagination.FormatResponse(err.Error(), nil, http.StatusBadRequest))
			c.Abort()
			return
		}
		msg = "Success Login"
		c.JSON(http.StatusOK, pagination.FormatResponse(msg, rs, http.StatusOK))

	}
}

func (a *authHadlerImpl) CurrentUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, _, err := token.ExtractTokenID(c)

		if err != nil {
			c.JSON(http.StatusBadRequest, pagination.FormatResponse(err.Error(), nil, http.StatusBadRequest))
			c.Abort()
			return
		}

		user, err := a.Service.CurrentUser(userId)
		if err != nil {
			c.JSON(http.StatusBadRequest, pagination.FormatResponse(err.Error(), nil, http.StatusBadRequest))
			c.Abort()
			return
		}

		msg = "Success Get Current User"
		c.JSON(http.StatusOK, pagination.FormatResponse(msg, user, http.StatusOK))

	}
}
