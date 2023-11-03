package main

import (
	"github.com/labstack/echo/v4"
	mysql "github.com/rizghz/genesys/infrastructure/database/MySql"
	"github.com/rizghz/genesys/internal/database/migration"
	"github.com/rizghz/genesys/internal/injector"
	"github.com/rizghz/genesys/internal/route"
	asisten "github.com/rizghz/genesys/module/Asisten"
	auth "github.com/rizghz/genesys/module/Auth"
	kelas "github.com/rizghz/genesys/module/Kelas"
	mahasiswa "github.com/rizghz/genesys/module/Mahasiswa"
	mata_praktikum "github.com/rizghz/genesys/module/MataPraktikum"
	penjadwalan "github.com/rizghz/genesys/module/Penjadwalan"
	praktikum "github.com/rizghz/genesys/module/Praktikum"
	ruangan "github.com/rizghz/genesys/module/Ruangan"
)

func main() {
	driver := mysql.NewMySqlDriver()
	migrate(driver)

	e := echo.New()
	route.Middleware(e)

	AuthModule(e, driver)
	UserModule(e, driver)
	MataPraktikumModule(e, driver)
	RuanganModule(e, driver)
	KelasModule(e, driver)
	AsistenModule(e, driver)
	MahasiswaModule(e, driver)
	JadwalModule(e, driver)
	PraktikumModule(e, driver)

	e.Logger.Fatal(e.Start(":8008"))
}

func migrate(driver *mysql.MySqlDriver) {
	migrator := migration.NewMySqlMigrator(driver)
	migrator.DropTable([]migration.Table{
		auth.UserModel{},
		auth.CredentialModel{},
		kelas.KelasModel{},
		mahasiswa.MahasiswaModel{},
		ruangan.RuanganModel{},
		asisten.AsistenModel{},
		mata_praktikum.MataPraktikumModel{},
		penjadwalan.PenjadwalanModel{},
		praktikum.PraktikumModel{},
	}...)
	migrator.CreateTable([]migration.Table{
		auth.UserModel{},
		auth.CredentialModel{},
		kelas.KelasModel{},
		mahasiswa.MahasiswaModel{},
		ruangan.RuanganModel{},
		asisten.AsistenModel{},
		mata_praktikum.MataPraktikumModel{},
		penjadwalan.PenjadwalanModel{},
		praktikum.PraktikumModel{},
	}...)
}

func AuthModule(e *echo.Echo, driver *mysql.MySqlDriver) {
	handler := injector.AuthInject(e, driver)
	route.Auth(e, handler)
}

func UserModule(e *echo.Echo, driver *mysql.MySqlDriver) {
	handler := injector.UserInject(e, driver)
	route.User(e, handler)
}

func MataPraktikumModule(e *echo.Echo, driver *mysql.MySqlDriver) {
	handler := injector.MatkumInject(e, driver)
	route.MataPraktikum(e, handler)
}

func RuanganModule(e *echo.Echo, driver *mysql.MySqlDriver) {
	handler := injector.RuanganInject(e, driver)
	route.Ruangan(e, handler)
}

func KelasModule(e *echo.Echo, driver *mysql.MySqlDriver) {
	handler := injector.KelasInject(e, driver)
	route.Kelas(e, handler)
}

func AsistenModule(e *echo.Echo, driver *mysql.MySqlDriver) {
	handler := injector.AsistenInject(e, driver)
	route.Asisten(e, handler)
}

func MahasiswaModule(e *echo.Echo, driver *mysql.MySqlDriver) {
	handler := injector.MahasiswaInject(e, driver)
	route.Mahasiswa(e, handler)
}

func JadwalModule(e *echo.Echo, driver *mysql.MySqlDriver) {
	handler := injector.PenjadwalanInject(e, driver)
	route.Jadwal(e, handler)
}

func PraktikumModule(e *echo.Echo, driver *mysql.MySqlDriver) {
	handler := injector.PraktikumInject(e, driver)
	route.Praktikum(e, handler)
}
