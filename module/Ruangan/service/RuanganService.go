package service

import (
	"github.com/labstack/echo/v4"
	"github.com/rizghz/genesys/module/Ruangan/transfer"
)

type RuanganService interface {
	GetSemuaRuangan(ctx echo.Context) []transfer.Response
	GetRuanganSpesifik(ctx echo.Context, kode string) *transfer.Response
	TambahRuangan(ctx echo.Context, request *transfer.RequestBody) *transfer.Response
	EditRuangan(ctx echo.Context, kode string, request *transfer.RequestBody) *transfer.Response
	HapusRuangan(ctx echo.Context, kode string) bool
}
