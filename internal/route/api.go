package route

import (
	echo "github.com/labstack/echo/v4"
	mid "github.com/rizghz/genesys/infrastructure/middleware"
	jwt "github.com/rizghz/genesys/infrastructure/middleware/JWT"
	asisten "github.com/rizghz/genesys/module/Asisten/handler"
	auth "github.com/rizghz/genesys/module/Auth/handler"
	kelas "github.com/rizghz/genesys/module/Kelas/handler"
	mahasiswa "github.com/rizghz/genesys/module/Mahasiswa/handler"
	matkum "github.com/rizghz/genesys/module/MataPraktikum/handler"
	ruangan "github.com/rizghz/genesys/module/Ruangan/handler"
)

func Middleware(e *echo.Echo) {
	e.Use([]echo.MiddlewareFunc{
		mid.Logger(),
		mid.Timeout(),
		mid.Cors(),
		mid.Uri(),
	}...)
}

func Auth(e *echo.Echo, handler auth.AuthHandler) {
	e.POST("/register", handler.Register())
	e.POST("/login", handler.Login())
	e.POST("/refresh", handler.Refresh())
	e.POST("/logout", handler.Logout(), jwt.Impl())
}

func User(e *echo.Echo, handler auth.UserHandler) {
	e.GET("/users", handler.Index(), jwt.Impl())
	e.GET("/users/:id", handler.Observe(), jwt.Impl())
	e.POST("/users", handler.Store(), jwt.Impl())
	e.PUT("/users/:id", handler.Edit(), jwt.Impl())
	e.DELETE("/users/:id", handler.Destroy(), jwt.Impl())
	// e.GET("/users/creds", nil, jwt.Impl())
	// e.GET("/users/:id/creds", nil, jwt.Impl())
}

func Kelas(e *echo.Echo, handler kelas.KelasHandler) {
	e.GET("/kelas", handler.Index(), jwt.Impl())
	e.GET("/kelas/:kode", handler.Observe(), jwt.Impl())
	e.POST("/kelas", handler.Store(), jwt.Impl())
	e.PUT("/kelas/:kode", handler.Edit(), jwt.Impl())
	e.DELETE("/kelas/:kode", handler.Destroy(), jwt.Impl())
}

func Ruangan(e *echo.Echo, handler ruangan.RuanganHandler) {
	e.GET("/ruangan", handler.Index(), jwt.Impl())
	e.GET("/ruangan/:kode", handler.Observe(), jwt.Impl())
	e.POST("/ruangan", handler.Store(), jwt.Impl())
	e.PUT("/ruangan/:kode", handler.Edit(), jwt.Impl())
	e.DELETE("/ruangan/:kode", handler.Destroy(), jwt.Impl())
}

func MataPraktikum(e *echo.Echo, handler matkum.MataPraktikumHandler) {
	e.GET("/matapraktikum", handler.Index(), jwt.Impl())
	e.GET("/matapraktikum/:kode", handler.Observe(), jwt.Impl())
	e.POST("/matapraktikum", handler.Store(), jwt.Impl())
	e.PUT("/matapraktikum/:kode", handler.Edit(), jwt.Impl())
	e.DELETE("/matapraktikum/:kode", handler.Destroy(), jwt.Impl())
}

func Mahasiswa(e *echo.Echo, handler mahasiswa.MahasiswaHandler) {
	e.GET("/mahasiswa", handler.Index(), jwt.Impl())
	e.GET("/mahasiswa/:npm", handler.Observe(), jwt.Impl())
	e.POST("/mahasiswa", handler.Store(), jwt.Impl())
	e.PUT("/mahasiswa/:npm", handler.Edit(), jwt.Impl())
	e.DELETE("/mahasiswa/:npm", handler.Destroy(), jwt.Impl())
}

func Asisten(e *echo.Echo, handler asisten.AsistenHandler) {
	e.GET("/asisten", handler.Index(), jwt.Impl())
	e.GET("/asisten/:nias", handler.Observe(), jwt.Impl())
	e.POST("/asisten", handler.Store(), jwt.Impl())
	e.PUT("/asisten/:nias", handler.Edit(), jwt.Impl())
	e.DELETE("/asisten/:nias", handler.Destroy(), jwt.Impl())
}
