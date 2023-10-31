package routes

import (
	"code-ai/handler"
	"code-ai/middleware"
	"code-ai/repository"
	"code-ai/services"
	"os"

	middlewares "github.com/labstack/echo/v4/middleware"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func AdminRouteInit(e *echo.Echo, DB *gorm.DB) {

	adminRepo := repository.NewAdminRepository(DB)
	adminService := services.NewAdminService(*adminRepo, &validator.Validate{})
	adminHandler := handler.NewAdminHandler(adminService)

	middleware.Logger(e)
	middleware.RateLimiter(e)
	middleware.Recover(e)
	middleware.CORS(e)
	middleware.NotFoundHandler(e)

	JWT := middlewares.JWT([]byte(os.Getenv("SECRET_KEY")))

	adminGroup := e.Group("/admin")
	adminGroup.POST("/regis/", adminHandler.RegisterAdminHandler)
	adminGroup.POST("/login/", adminHandler.LoginAdminHandler)
	adminGroup.GET("/", adminHandler.FindAllAdminHandler, JWT)
	adminGroup.GET("/:email", adminHandler.FindAdminByEmailHandler, JWT)
	adminGroup.GET("/user", adminHandler.FindAllUserHandler, JWT)
	adminGroup.GET("/user/:username", adminHandler.FindUserByUsernameHandler, JWT)
	adminGroup.PUT("/user/:id", adminHandler.UpdateUserHandler, JWT)
	adminGroup.DELETE("/user/:id", adminHandler.DeleteUserHandler, JWT)
}