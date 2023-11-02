package main

import (
	"github.com/labstack/echo/v4"
	mysql "github.com/rizghz/genesys/infrastructure/database/MySql"
	"github.com/rizghz/genesys/infrastructure/service/cloudinary"
	"github.com/rizghz/genesys/internal/database/migration"
	"github.com/rizghz/genesys/internal/route"
	asisten "github.com/rizghz/genesys/module/Asisten"
	be "github.com/rizghz/genesys/module/Auth"
	bh "github.com/rizghz/genesys/module/Auth/handler"
	br "github.com/rizghz/genesys/module/Auth/repository"
	bs "github.com/rizghz/genesys/module/Auth/service"
	ke "github.com/rizghz/genesys/module/Kelas"
	kh "github.com/rizghz/genesys/module/Kelas/handler"
	kr "github.com/rizghz/genesys/module/Kelas/repository"
	ks "github.com/rizghz/genesys/module/Kelas/service"
	ae "github.com/rizghz/genesys/module/Mahasiswa"
	me "github.com/rizghz/genesys/module/MataPraktikum"
	mh "github.com/rizghz/genesys/module/MataPraktikum/handler"
	mr "github.com/rizghz/genesys/module/MataPraktikum/repository"
	ms "github.com/rizghz/genesys/module/MataPraktikum/service"
	penjadwalan "github.com/rizghz/genesys/module/Penjadwalan"
	praktikum "github.com/rizghz/genesys/module/Praktikum"
	re "github.com/rizghz/genesys/module/Ruangan"
	rh "github.com/rizghz/genesys/module/Ruangan/handler"
	rr "github.com/rizghz/genesys/module/Ruangan/repository"
	rs "github.com/rizghz/genesys/module/Ruangan/service"
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

	e.Logger.Fatal(e.Start(":8008"))
}

func migrate(driver *mysql.MySqlDriver) {
	migrator := migration.NewMySqlMigrator(driver)
	migrator.DropTable([]migration.Table{
		asisten.AsistenModel{},
		penjadwalan.PenjadwalanModel{},
		be.UserModel{},
		be.CredentialModel{},
		ke.KelasModel{},
		ae.MahasiswaModel{},
		re.RuanganModel{},
		me.MataPraktikumModel{},
		praktikum.PraktikumModel{},
	}...)
	migrator.CreateTable([]migration.Table{
		asisten.AsistenModel{},
		penjadwalan.PenjadwalanModel{},
		be.UserModel{},
		be.CredentialModel{},
		ke.KelasModel{},
		ae.MahasiswaModel{},
		re.RuanganModel{},
		me.MataPraktikumModel{},
		praktikum.PraktikumModel{},
	}...)
}

func UserModule(e *echo.Echo, driver *mysql.MySqlDriver) {
	repository := br.NewUserMySqlRepository(driver)
	service := bs.NewUserServiceImpl(repository, cloudinary.NewImageUploader())
	handler := bh.NewUserHttpHandler(service)
	route.User(e, handler)
}

func AuthModule(e *echo.Echo, driver *mysql.MySqlDriver) {
	repository := br.NewCredentialMySqlRepository(driver)
	service := bs.NewAuthServiceImpl(repository)
	handler := bh.NewAuthHttpHandler(service)
	route.Auth(e, handler)
}

func KelasModule(e *echo.Echo, driver *mysql.MySqlDriver) {
	repository := kr.NewKelasMySqlRepository(driver)
	service := ks.NewKelasServiceImpl(repository)
	handler := kh.NewKelasHttpHandler(service)
	route.Kelas(e, handler)
}

func RuanganModule(e *echo.Echo, driver *mysql.MySqlDriver) {
	repository := rr.NewRuanganMySqlRepository(driver)
	service := rs.NewRuanganServiceImpl(repository)
	handler := rh.NewRuanganHttpHandler(service)
	route.Ruangan(e, handler)
}

func MataPraktikumModule(e *echo.Echo, driver *mysql.MySqlDriver) {
	repository := mr.NewMataPraktikumMySqlRepository(driver)
	service := ms.NewMataPraktikumServiceImpl(repository)
	handler := mh.NewMataPraktikumHttpHandler(service)
	route.MataPraktikum(e, handler)
}
