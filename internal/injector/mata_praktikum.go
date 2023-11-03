package injector

import (
	"github.com/labstack/echo/v4"
	mysql "github.com/rizghz/genesys/infrastructure/database/MySql"
	"github.com/rizghz/genesys/module/MataPraktikum/handler"
	"github.com/rizghz/genesys/module/MataPraktikum/repository"
	"github.com/rizghz/genesys/module/MataPraktikum/service"
)

func MatkumInject(e *echo.Echo, driver *mysql.MySqlDriver) handler.MatkumHandler {
	repository := repository.NewMatkumMySqlRepository(driver)
	service := service.NewMatkumServiceImpl(repository)
	handler := handler.NewMatkumHttpHandler(service)
	return handler
}
