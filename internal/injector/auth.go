package injector

import (
	"github.com/labstack/echo/v4"
	mysql "github.com/rizghz/genesys/infrastructure/database/MySql"
	"github.com/rizghz/genesys/infrastructure/service/cloudinary"
	"github.com/rizghz/genesys/module/Auth/handler"
	"github.com/rizghz/genesys/module/Auth/repository"
	"github.com/rizghz/genesys/module/Auth/service"
)

func AuthInject(e *echo.Echo, driver *mysql.MySqlDriver) handler.AuthHandler {
	repository := repository.NewCredentialMySqlRepository(driver)
	service := service.NewAuthServiceImpl(repository)
	handler := handler.NewAuthHttpHandler(service)
	return handler
}

func UserInject(e *echo.Echo, driver *mysql.MySqlDriver) handler.UserHandler {
	repository := repository.NewUserMySqlRepository(driver)
	service := service.NewUserServiceImpl(repository, cloudinary.NewImageUploader())
	handler := handler.NewUserHttpHandler(service)
	return handler
}
