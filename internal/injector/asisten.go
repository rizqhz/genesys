package injector

import (
	"github.com/labstack/echo/v4"
	mysql "github.com/rizghz/genesys/infrastructure/database/MySql"
	"github.com/rizghz/genesys/module/Asisten/handler"
	"github.com/rizghz/genesys/module/Asisten/repository"
	"github.com/rizghz/genesys/module/Asisten/service"
)

func AsistenInject(e *echo.Echo, driver *mysql.MySqlDriver) handler.AsistenHandler {
	repository := repository.NewAsistenMySqlRepository(driver)
	service := service.NewAsistenServiceImpl(repository)
	handler := handler.NewAsistenHttpHandler(service)
	return handler
}
