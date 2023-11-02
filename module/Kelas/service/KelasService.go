package service

import (
	"net/url"

	"github.com/rizghz/genesys/module/Kelas/transfer"
)

type KelasService interface {
	GetSemuaKelas(query url.Values) []transfer.Response
	GetKelasSpesifik(kode string) *transfer.Response
	TambahKelas(data transfer.RequestBody) *transfer.Response
	EditKelas(kode string, data transfer.RequestBody) *transfer.Response
	HapusKelas(kode string) bool
}
