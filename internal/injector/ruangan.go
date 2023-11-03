package injector

import (
	"github.com/labstack/echo/v4"
	mysql "github.com/rizghz/genesys/infrastructure/database/MySql"
	"github.com/rizghz/genesys/module/Ruangan/handler"
	"github.com/rizghz/genesys/module/Ruangan/repository"
	"github.com/rizghz/genesys/module/Ruangan/service"
)

func RuanganInject(e *echo.Echo, driver *mysql.MySqlDriver) handler.RuanganHandler {
	repository := repository.NewRuanganMySqlRepository(driver)
	service := service.NewRuanganServiceImpl(repository)
	handler := handler.NewRuanganHttpHandler(service)
	return handler
}
