package injector

import (
	"github.com/labstack/echo/v4"
	mysql "github.com/rizghz/genesys/infrastructure/database/MySql"
	"github.com/rizghz/genesys/module/Mahasiswa/handler"
	"github.com/rizghz/genesys/module/Mahasiswa/repository"
	"github.com/rizghz/genesys/module/Mahasiswa/service"
)

func MahasiswaInject(e *echo.Echo, driver *mysql.MySqlDriver) handler.MahasiswaHandler {
	repository := repository.NewMahasiswaMySqlRepository(driver)
	service := service.NewMahasiswaServiceImpl(repository)
	handler := handler.NewMahasiswaHttpHandler(service)
	return handler
}
