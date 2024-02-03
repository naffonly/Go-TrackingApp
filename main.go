package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-playground/validator/v10"
	configure "trackingApp/config"
	handler2 "trackingApp/features/geoapify/handler"
	"trackingApp/features/vehicle/handler"
	"trackingApp/features/vehicle/repository"
	"trackingApp/features/vehicle/service"
	"trackingApp/middleware"
	"trackingApp/routes"
	"trackingApp/utils/database"

	locationHandler "trackingApp/features/location/handler"
	locationRepository "trackingApp/features/location/repository"
	locationService "trackingApp/features/location/service"

	companyHandler "trackingApp/features/company/handler"
	companyRepository "trackingApp/features/company/repository"
	companyService "trackingApp/features/company/service"

	autHandler "trackingApp/features/auth/handler"
	authRepository "trackingApp/features/auth/repository"
	authService "trackingApp/features/auth/service"

	userHandler "trackingApp/features/user/handler"
	userRepository "trackingApp/features/user/repository"
	userService "trackingApp/features/user/service"

	orderHandler "trackingApp/features/order/handler"
	orderRepository "trackingApp/features/order/repository"
	orderService "trackingApp/features/order/service"
)

func main() {
	SetupAppRouter()
}

func SetupAppRouter() *gin.Engine {
	config := configure.InitConfig()
	router := gin.Default()

	validation := validator.New()
	db := database.InitDB(config)

	public := router.Group("/api/v1")
	protected := router.Group("/api/v1")
	protected.Use(middleware.AuthValid)
	//Auth
	authRepo := authRepository.NewAuthRepository(db)
	authSvc := authService.NewAuthService(authRepo, validation)
	authHdlr := autHandler.NewAuthHandler(authSvc)
	//User
	userRepo := userRepository.NewUserRepositoryInterface(db)
	userSvc := userService.NewUserServiceInterface(userRepo, validation)
	userHdlr := userHandler.NewUserHandlerInterface(userSvc)
	//Company
	companyRepo := companyRepository.NewCompanyRepositoryImpl(db)
	companySvc := companyService.NewCompanyServiceInterface(companyRepo, validation)
	companyHdlr := companyHandler.NewCompanyHanlderImpl(companySvc)
	//Location
	locationRepo := locationRepository.NewLocationRepositoryImpl(db)
	locationSvc := locationService.NewLocationSeriveImpl(locationRepo, validation)
	locationHdlr := locationHandler.NewLocationHandlerInterface(locationSvc)
	//Order
	orderRepo := orderRepository.NewOrderRepositoryImpl(db)
	orderSvc := orderService.NewOrderServiceImpl(orderRepo, validation)
	orderHdl := orderHandler.NewOrderHandlerImpl(orderSvc)
	//Vehicle
	vehicleRepo := repository.NewVehicleRepositoryImpl(db)
	vehicleSvc := service.NewVehicleServiceImpl(vehicleRepo, validation)
	vehicleHhl := handler.NewVehicleHandlerImpl(vehicleSvc)
	//Geo
	geoHandler := handler2.NewGeoapify(db, config.GeoKey)

	routes.InitRoutesPublic(public, authHdlr, geoHandler)
	routes.InitRoutesPrivate(protected, userHdlr, authHdlr, companyHdlr, locationHdlr, orderHdl, vehicleHhl)

	router.Run(config.SeverPort)
	return router
}
