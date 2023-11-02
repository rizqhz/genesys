package main

import (
	"github.com/labstack/echo/v4"
	mysql "github.com/rizghz/genesys/infrastructure/database/MySql"
	"github.com/rizghz/genesys/infrastructure/service/cloudinary"
	"github.com/rizghz/genesys/internal/database/migration"
	"github.com/rizghz/genesys/internal/route"
	. "github.com/rizghz/genesys/module/Asisten"
	. "github.com/rizghz/genesys/module/Asisten/handler"
	. "github.com/rizghz/genesys/module/Asisten/repository"
	. "github.com/rizghz/genesys/module/Asisten/service"
	. "github.com/rizghz/genesys/module/Auth"
	. "github.com/rizghz/genesys/module/Auth/handler"
	auth_r "github.com/rizghz/genesys/module/Auth/repository"
	. "github.com/rizghz/genesys/module/Auth/service"
	. "github.com/rizghz/genesys/module/Kelas"
	. "github.com/rizghz/genesys/module/Kelas/handler"
	kelas_r "github.com/rizghz/genesys/module/Kelas/repository"
	. "github.com/rizghz/genesys/module/Kelas/service"
	. "github.com/rizghz/genesys/module/Mahasiswa"
	. "github.com/rizghz/genesys/module/Mahasiswa/handler"
	mahasiswa_r "github.com/rizghz/genesys/module/Mahasiswa/repository"
	. "github.com/rizghz/genesys/module/Mahasiswa/service"
	. "github.com/rizghz/genesys/module/MataPraktikum"
	. "github.com/rizghz/genesys/module/MataPraktikum/handler"
	mata_r "github.com/rizghz/genesys/module/MataPraktikum/repository"
	. "github.com/rizghz/genesys/module/MataPraktikum/service"
	. "github.com/rizghz/genesys/module/Penjadwalan"
	. "github.com/rizghz/genesys/module/Praktikum"
	. "github.com/rizghz/genesys/module/Ruangan"
	. "github.com/rizghz/genesys/module/Ruangan/handler"
	ruang_r "github.com/rizghz/genesys/module/Ruangan/repository"
	. "github.com/rizghz/genesys/module/Ruangan/service"
)

func main() {
	driver := mysql.NewMySqlDriver()
	migrate(driver)

	e := echo.New()
	route.Middleware(e)

	UserModule(e, driver)
	AuthModule(e, driver)
	KelasModule(e, driver)
	RuanganModule(e, driver)
	MataPraktikumModule(e, driver)
	MahasiswaModule(e, driver)
	AsistenModule(e, driver)

	e.Logger.Fatal(e.Start(":8008"))
}

func migrate(driver *mysql.MySqlDriver) {
	migrator := migration.NewMySqlMigrator(driver)
	migrator.DropTable([]migration.Table{
		UserModel{},
		CredentialModel{},
		KelasModel{},
		MahasiswaModel{},
		RuanganModel{},
		AsistenModel{},
		MataPraktikumModel{},
		PenjadwalanModel{},
		PraktikumModel{},
	}...)
	migrator.CreateTable([]migration.Table{
		UserModel{},
		CredentialModel{},
		KelasModel{},
		MahasiswaModel{},
		RuanganModel{},
		AsistenModel{},
		MataPraktikumModel{},
		PenjadwalanModel{},
		PraktikumModel{},
	}...)
}

func UserModule(e *echo.Echo, driver *mysql.MySqlDriver) {
	repository := auth_r.NewUserMySqlRepository(driver)
	service := NewUserServiceImpl(repository, cloudinary.NewImageUploader())
	handler := NewUserHttpHandler(service)
	route.User(e, handler)
}

func AuthModule(e *echo.Echo, driver *mysql.MySqlDriver) {
	repository := auth_r.NewCredentialMySqlRepository(driver)
	service := NewAuthServiceImpl(repository)
	handler := NewAuthHttpHandler(service)
	route.Auth(e, handler)
}

func KelasModule(e *echo.Echo, driver *mysql.MySqlDriver) {
	repository := kelas_r.NewKelasMySqlRepository(driver)
	service := NewKelasServiceImpl(repository)
	handler := NewKelasHttpHandler(service)
	route.Kelas(e, handler)
}

func RuanganModule(e *echo.Echo, driver *mysql.MySqlDriver) {
	repository := ruang_r.NewRuanganMySqlRepository(driver)
	service := NewRuanganServiceImpl(repository)
	handler := NewRuanganHttpHandler(service)
	route.Ruangan(e, handler)
}

func MataPraktikumModule(e *echo.Echo, driver *mysql.MySqlDriver) {
	repository := mata_r.NewMataPraktikumMySqlRepository(driver)
	service := NewMataPraktikumServiceImpl(repository)
	handler := NewMataPraktikumHttpHandler(service)
	route.MataPraktikum(e, handler)
}

func MahasiswaModule(e *echo.Echo, driver *mysql.MySqlDriver) {
	repository := mahasiswa_r.NewMahasiswaMySqlRepository(driver)
	service := NewMahasiswaServiceImpl(repository)
	handler := NewMahasiswaHttpHandler(service)
	route.Mahasiswa(e, handler)
}

func AsistenModule(e *echo.Echo, driver *mysql.MySqlDriver) {
	repository := NewAsistenMySqlRepository(driver)
	service := NewAsistenServiceImpl(repository)
	handler := NewAsistenHttpHandler(service)
	route.Asisten(e, handler)
}
