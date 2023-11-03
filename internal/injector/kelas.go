package injector

import (
	"github.com/labstack/echo/v4"
	mysql "github.com/rizghz/genesys/infrastructure/database/MySql"
	"github.com/rizghz/genesys/module/Kelas/handler"
	"github.com/rizghz/genesys/module/Kelas/repository"
	"github.com/rizghz/genesys/module/Kelas/service"
)

func KelasInject(e *echo.Echo, driver *mysql.MySqlDriver) handler.KelasHandler {
	repository := repository.NewKelasMySqlRepository(driver)
	service := service.NewKelasServiceImpl(repository)
	handler := handler.NewKelasHttpHandler(service)
	return handler
}
