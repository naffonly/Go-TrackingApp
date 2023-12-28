package main

import (
	"github.com/gin-gonic/gin"
	configure "trackingApp/config"
	"trackingApp/middleware"

	companyHandler "trackingApp/features/company/handler"
	companyRepository "trackingApp/features/company/repository"
	companyService "trackingApp/features/company/service"

	autHandler "trackingApp/features/auth/handler"
	authRepository "trackingApp/features/auth/repository"
	authService "trackingApp/features/auth/service"

	userHandler "trackingApp/features/user/handler"
	userRepository "trackingApp/features/user/repository"
	userService "trackingApp/features/user/service"
	"trackingApp/routes"
	"trackingApp/utils/database"
)

func main() {
	SetupAppRouter()
}

func SetupAppRouter() *gin.Engine {
	config := configure.InitConfig()
	router := gin.Default()

	db := database.InitDB(config)

	public := router.Group("/api/v1")
	protected := router.Group("/api/v1")
	protected.Use(middleware.AuthValid)

	authRepo := authRepository.NewAuthRepository(db)
	authSvc := authService.NewAuthService(authRepo)
	authHdlr := autHandler.NewAuthHandler(authSvc)

	userRepo := userRepository.NewUserRepositoryInterface(db)
	userSvc := userService.NewUserServiceInterface(userRepo)
	userHdlr := userHandler.NewUserHandlerInterface(userSvc)

	companyRepo := companyRepository.NewCompanyRepositoryImpl(db)
	companySvc := companyService.NewCompanyServiceInterface(companyRepo)
	companyHdlr := companyHandler.NewCompanyHanlderImpl(companySvc)

	routes.InitRoutesPublic(public, authHdlr)
	routes.InitRoutesPrivate(protected, userHdlr, authHdlr, companyHdlr)

	router.Run(config.SeverPort)
	return router
}
