package route

import (
	echo "github.com/labstack/echo/v4"
	mid "github.com/rizghz/genesys/infrastructure/middleware"
	jwt "github.com/rizghz/genesys/infrastructure/middleware/JWT"
	auth "github.com/rizghz/genesys/module/Auth/handler"
	kelas "github.com/rizghz/genesys/module/Kelas/handler"
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
	e.GET("/users", handler.Index())
	e.GET("/users/:id", handler.Observe())
	e.POST("/users", handler.Store())
	e.PUT("/users/:id", handler.Edit())
	e.DELETE("/users/:id", handler.Destroy())
	e.GET("/users/creds", nil)
	e.GET("/users/:id/creds", nil)
}

func Kelas(e *echo.Echo, handler kelas.KelasHandler) {
	e.GET("/kelas", handler.Index(), jwt.Impl())
	e.GET("/kelas/:kode", handler.Observe(), jwt.Impl())
	e.POST("/kelas", handler.Store(), jwt.Impl())
	e.PUT("/kelas/:kode", handler.Edit(), jwt.Impl())
	e.DELETE("/kelas/:kode", handler.Destroy(), jwt.Impl())
}

func Ruangan(e *echo.Echo, handler ruangan.RuanganHandler) {
	e.GET("/ruangan", handler.Index())
	e.GET("/ruangan/:kode", handler.Observe())
	e.POST("/ruangan", handler.Store())
	e.PUT("/ruangan/:kode", handler.Edit())
	e.DELETE("/ruangan/:kode", handler.Destroy())
}

func MataPraktikum(e *echo.Echo, handler matkum.MataPraktikumHandler) {
	e.GET("/matapraktikum", handler.Index())
	e.GET("/matapraktikum/:kode", handler.Observe())
	e.POST("/matapraktikum", handler.Store())
	e.PUT("/matapraktikum/:kode", handler.Edit())
	e.DELETE("/matapraktikum/:kode", handler.Destroy())
}
