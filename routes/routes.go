package routes

import (
	"github.com/gin-gonic/gin"
	autHandler "trackingApp/features/auth/handler"
	companyHandler "trackingApp/features/company/handler"
	geoHandler "trackingApp/features/geoapify/handler"
	locationHandler "trackingApp/features/location/handler"
	orderHandler "trackingApp/features/order/handler"
	"trackingApp/features/user/handler"
	vehicleHandler "trackingApp/features/vehicle/handler"
)

func InitRoutesPublic(r *gin.RouterGroup, autHandler autHandler.AuthHandlerInterface, geoHandler geoHandler.GeoapifyInterface) {
	authRoutes(r, autHandler)
	geoRoutes(r, geoHandler)
}

func InitRoutesPrivate(r *gin.RouterGroup,
	user handler.UserHandlerInterface,
	auth autHandler.AuthHandlerInterface,
	company companyHandler.CompanyHandlerInterface,
	location locationHandler.LocationHandlerInterface,
	order orderHandler.OrderHandlerInterface,
	vehicle vehicleHandler.VehicleHandlerInterface,
) {
	userRoutes(r, user)
	profilRoutes(r, auth)
	companyRoutes(r, company)
	locationRoutes(r, location)
	orderRoutes(r, order)
	vehicleRoutes(r, vehicle)
}

func userRoutes(router *gin.RouterGroup, handlerInterface handler.UserHandlerInterface) {
	router.GET("/user", handlerInterface.FindAll())
	router.GET("/user/:id", handlerInterface.FindById())
	router.POST("/user", handlerInterface.Insert())
	router.PUT("/user/:id", handlerInterface.Update())
	router.DELETE("/user/:id", handlerInterface.Delete())
	router.GET("/user-company", handlerInterface.GetCurrentCompany())
}

func authRoutes(router *gin.RouterGroup, handlerInterface autHandler.AuthHandlerInterface) {
	router.POST("/login", handlerInterface.Login())
}

func geoRoutes(router *gin.RouterGroup, handlerInterface geoHandler.GeoapifyInterface) {
	router.POST("/geo", handlerInterface.Insert())
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

func locationRoutes(router *gin.RouterGroup, handlerInterface locationHandler.LocationHandlerInterface) {
	router.GET("/location", handlerInterface.FindAll())
	router.GET("/location/:id", handlerInterface.FindByID())
	router.POST("/location", handlerInterface.Insert())
	router.PUT("/location/:id", handlerInterface.Update())
	router.DELETE("/location/:id", handlerInterface.Delete())
}

func orderRoutes(router *gin.RouterGroup, handlerInterface orderHandler.OrderHandlerInterface) {
	router.GET("/order", handlerInterface.FindAll())
	router.GET("/order/:id", handlerInterface.FindByID())
	router.GET("/order-identity/:identity", handlerInterface.GetIdentity())
	router.POST("/order", handlerInterface.Insert())
	router.PUT("/order/:id", handlerInterface.Update())
	router.DELETE("/order/:id", handlerInterface.Delete())
}
func vehicleRoutes(router *gin.RouterGroup, handlerInterface vehicleHandler.VehicleHandlerInterface) {
	router.GET("/vehicle", handlerInterface.FindAll())
	router.GET("/vehicle/:id", handlerInterface.FindByID())
	router.POST("/vehicle", handlerInterface.Insert())
	router.PUT("/vehicle/:id", handlerInterface.Update())
	router.DELETE("/vehicle/:id", handlerInterface.Delete())
}
