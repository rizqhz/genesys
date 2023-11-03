package service

import (
	"github.com/labstack/echo/v4"
	"github.com/rizghz/genesys/module/Mahasiswa/transfer"
)

type MahasiswaService interface {
	GetSemuaMahasiswa(ctx echo.Context) []transfer.Response
	GetMahasiswaSpesifik(ctx echo.Context, npm string) *transfer.Response
	TambahMahasiswa(ctx echo.Context, request *transfer.RequestBody) *transfer.Response
	EditMahasiswa(ctx echo.Context, npm string, request *transfer.RequestBody) *transfer.Response
	HapusMahasiswa(ctx echo.Context, npm string) bool
}
