package injector

import (
	"github.com/labstack/echo/v4"
	mysql "github.com/rizghz/genesys/infrastructure/database/MySql"
	"github.com/rizghz/genesys/module/Penjadwalan/handler"
	"github.com/rizghz/genesys/module/Penjadwalan/repository"
	"github.com/rizghz/genesys/module/Penjadwalan/service"
)

func PenjadwalanInject(e *echo.Echo, driver *mysql.MySqlDriver) handler.PenjadwalanHandler {
	repository := repository.NewPenjadwalanMySqlRepository(driver)
	service := service.NewPenjadwalanServiceImpl(repository)
	handler := handler.NewPenjadwalanHttpHandler(service)
	return handler
}
