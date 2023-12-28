package routes

import (
	"github.com/gin-gonic/gin"
	autHandler "trackingApp/features/auth/handler"
	companyHandler "trackingApp/features/company/handler"
	"trackingApp/features/user/handler"
)

func InitRoutesPublic(r *gin.RouterGroup, autHandler autHandler.AuthHandlerInterface) {
	authRoutes(r, autHandler)
}

func InitRoutesPrivate(r *gin.RouterGroup, user handler.UserHandlerInterface, auth autHandler.AuthHandlerInterface, company companyHandler.CompanyHandlerInterface) {
	userRoutes(r, user)
	profilRoutes(r, auth)
	companyRoutes(r, company)
}

func userRoutes(router *gin.RouterGroup, handlerInterface handler.UserHandlerInterface) {
	router.GET("/user", handlerInterface.FindAll())
	router.GET("/user/:id", handlerInterface.FindById())
	router.POST("/user", handlerInterface.Insert())
	router.PUT("/user/:id", handlerInterface.Update())
	router.DELETE("/user/:id", handlerInterface.Delete())
}

func authRoutes(router *gin.RouterGroup, handlerInterface autHandler.AuthHandlerInterface) {
	router.POST("/login", handlerInterface.Login())
}

func profilRoutes(router *gin.RouterGroup, handlerInterface autHandler.AuthHandlerInterface) {
	router.GET("/me", handlerInterface.CurrentUser())
}

func companyRoutes(router *gin.RouterGroup, handlerInterface companyHandler.CompanyHandlerInterface) {
	router.GET("/company", handlerInterface.FindAll())
	router.GET("/company/:id", handlerInterface.FindByID())
	router.POST("/company", handlerInterface.Insert())
	router.PUT("/company/:id", handlerInterface.Update())
	router.DELETE("/company/:id", handlerInterface.Delete())

}
