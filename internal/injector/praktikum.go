package injector

import (
	"github.com/labstack/echo/v4"
	mysql "github.com/rizghz/genesys/infrastructure/database/MySql"
	"github.com/rizghz/genesys/module/Praktikum/handler"
	"github.com/rizghz/genesys/module/Praktikum/repository"
	"github.com/rizghz/genesys/module/Praktikum/service"
)

func PraktikumInject(e *echo.Echo, driver *mysql.MySqlDriver) handler.PraktikumHandler {
	repository := repository.NewPraktikumMySqlRepository(driver)
	service := service.NewPraktikumServiceImpl(repository)
	handler := handler.NewPraktikumHttpHandler(service)
	return handler
}
