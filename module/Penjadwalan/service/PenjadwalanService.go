package service

import (
	"github.com/labstack/echo/v4"
	"github.com/rizghz/genesys/module/Penjadwalan/transfer"
)

type PenjadwalanService interface {
	GetSemuaJadwal(ctx echo.Context) []transfer.Response
	GetJadwalSpesifik(id int) *transfer.Response
	TambahJadwal(request *transfer.RequestBody) *transfer.Response
	EditJadwal(id int, request *transfer.RequestBody) *transfer.Response
	HapusJadwal(id int) bool
}
