package service

import (
	"github.com/labstack/echo/v4"
	"github.com/rizghz/genesys/module/Kelas/transfer"
)

type KelasService interface {
	GetSemuaKelas(ctx echo.Context) []transfer.Response
	GetKelasSpesifik(ctx echo.Context, kode string) *transfer.Response
	TambahKelas(ctx echo.Context, request *transfer.RequestBody) *transfer.Response
	EditKelas(ctx echo.Context, kode string, request *transfer.RequestBody) *transfer.Response
	HapusKelas(ctx echo.Context, kode string) bool
}
