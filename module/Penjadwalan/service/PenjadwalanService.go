package service

import (
	"github.com/labstack/echo/v4"
	"github.com/rizghz/genesys/module/Penjadwalan/transfer"
)

type PenjadwalanService interface {
	GetSemuaJadwal(ctx echo.Context) []transfer.Response
	GetJadwalSpesifik(ctx echo.Context, id int) *transfer.Response
	TambahJadwal(ctx echo.Context, request *transfer.RequestBody) *transfer.Response
	EditJadwal(ctx echo.Context, id int, request *transfer.RequestBody) *transfer.Response
	HapusJadwal(ctx echo.Context, id int) bool
}
