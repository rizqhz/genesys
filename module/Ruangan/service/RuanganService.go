package service

import (
	"net/url"

	"github.com/rizghz/genesys/module/Ruangan/transfer"
)

type RuanganService interface {
	GetSemuaRuangan(query url.Values) []transfer.Response
	GetRuanganSpesifik(kode string) *transfer.Response
	TambahRuangan(data transfer.RequestBody) *transfer.Response
	EditRuangan(kode string, data transfer.RequestBody) *transfer.Response
	HapusRuangan(kode string) bool
}
