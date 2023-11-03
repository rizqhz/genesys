package route

import (
	echo "github.com/labstack/echo/v4"
	"github.com/rizghz/genesys/infrastructure/middleware"
	jwt "github.com/rizghz/genesys/infrastructure/middleware/JWT"
	. "github.com/rizghz/genesys/module/Asisten/handler"
	. "github.com/rizghz/genesys/module/Auth/handler"
	. "github.com/rizghz/genesys/module/Kelas/handler"
	. "github.com/rizghz/genesys/module/Mahasiswa/handler"
	. "github.com/rizghz/genesys/module/MataPraktikum/handler"
	. "github.com/rizghz/genesys/module/Penjadwalan/handler"
	. "github.com/rizghz/genesys/module/Praktikum/handler"
	. "github.com/rizghz/genesys/module/Ruangan/handler"
)

func Middleware(e *echo.Echo) {
	e.Use([]echo.MiddlewareFunc{
		middleware.Logger(),
		middleware.Timeout(),
		middleware.Cors(),
		middleware.Uri(),
	}...)
}

func Auth(e *echo.Echo, handler AuthHandler) {
	e.POST("/register", handler.Register()) // --- // ✓
	e.POST("/login", handler.Login())       // --- // ✓
	e.POST("/refresh", handler.Refresh(), jwt.Impl())
	e.POST("/logout", handler.Logout(), jwt.Impl())
}

func User(e *echo.Echo, handler UserHandler) {
	user := e.Group("/users")
	user.GET("", handler.Index(), jwt.Impl())          // --- // ✓
	user.GET("/:id", handler.Observe(), jwt.Impl())    // --- // ✓
	user.POST("", handler.Store(), jwt.Impl())         // --- // ✓
	user.PUT("/:id", handler.Edit(), jwt.Impl())       // --- // ✓
	user.DELETE("/:id", handler.Destroy(), jwt.Impl()) // --- // ✓
}

func MataPraktikum(e *echo.Echo, handler MatkumHandler) {
	matkum := e.Group("/matkum")
	matkum.GET("", handler.Index(), jwt.Impl())            // --- // ✓
	matkum.GET("/:kode", handler.Observe(), jwt.Impl())    // --- // ✓
	matkum.POST("", handler.Store(), jwt.Impl())           // --- // ✓
	matkum.PUT("/:kode", handler.Edit(), jwt.Impl())       // --- // ✓
	matkum.DELETE("/:kode", handler.Destroy(), jwt.Impl()) // --- // ✓
}

func Ruangan(e *echo.Echo, handler RuanganHandler) {
	ruangan := e.Group("/ruangan")
	ruangan.GET("", handler.Index(), jwt.Impl())            // --- // ✓
	ruangan.GET("/:kode", handler.Observe(), jwt.Impl())    // --- // ✓
	ruangan.POST("", handler.Store(), jwt.Impl())           // --- // ✓
	ruangan.PUT("/:kode", handler.Edit(), jwt.Impl())       // --- // ✓
	ruangan.DELETE("/:kode", handler.Destroy(), jwt.Impl()) // --- // ✓
}

func Kelas(e *echo.Echo, handler KelasHandler) {
	kelas := e.Group("/kelas")
	kelas.GET("", handler.Index(), jwt.Impl())            // --- // ✓
	kelas.GET("/:kode", handler.Observe(), jwt.Impl())    // --- // ✓
	kelas.POST("", handler.Store(), jwt.Impl())           // --- // ✓
	kelas.PUT("/:kode", handler.Edit(), jwt.Impl())       // --- // ✓
	kelas.DELETE("/:kode", handler.Destroy(), jwt.Impl()) // --- // ✓
}

func Asisten(e *echo.Echo, handler AsistenHandler) {
	asisten := e.Group("/asisten")
	asisten.GET("", handler.Index(), jwt.Impl())            // --- // ✓
	asisten.GET("/:nias", handler.Observe(), jwt.Impl())    // --- // ✓
	asisten.POST("", handler.Store(), jwt.Impl())           // --- // ✓
	asisten.PUT("/:nias", handler.Edit(), jwt.Impl())       // --- // ✓
	asisten.DELETE("/:nias", handler.Destroy(), jwt.Impl()) // --- // ✓
}

func Mahasiswa(e *echo.Echo, handler MahasiswaHandler) {
	mahasiswa := e.Group("/mahasiswa")
	mahasiswa.GET("", handler.Index(), jwt.Impl())           // --- // ✓
	mahasiswa.GET("/:npm", handler.Observe(), jwt.Impl())    // --- // ✓
	mahasiswa.POST("", handler.Store(), jwt.Impl())          // --- // ✓
	mahasiswa.PUT("/:npm", handler.Edit(), jwt.Impl())       // --- // ✓
	mahasiswa.DELETE("/:npm", handler.Destroy(), jwt.Impl()) // --- // ✓
}

func Jadwal(e *echo.Echo, handler PenjadwalanHandler) {
	jadwal := e.Group("/jadwal")
	jadwal.GET("", handler.Index(), jwt.Impl())          // --- // ✓
	jadwal.GET("/:id", handler.Observe(), jwt.Impl())    // --- // ✓
	jadwal.POST("", handler.Store(), jwt.Impl())         // --- // ✓
	jadwal.PUT("/:id", handler.Edit(), jwt.Impl())       // --- // ✓
	jadwal.DELETE("/:id", handler.Destroy(), jwt.Impl()) // --- // ✓
}

func Praktikum(e *echo.Echo, handler PraktikumHandler) {
	praktikum := e.Group("/praktikum")
	praktikum.GET("", handler.Index(), jwt.Impl())          // --- // ✓
	praktikum.GET("/:id", handler.Observe(), jwt.Impl())    // --- // ✓
	praktikum.POST("", handler.Store(), jwt.Impl())         // --- // ✓
	praktikum.PUT("/:id", handler.Edit(), jwt.Impl())       // --- // ✓
	praktikum.DELETE("/:id", handler.Destroy(), jwt.Impl()) // --- // ✓
	// praktikum.GET("/generate", nil, jwt.Impl())
}
