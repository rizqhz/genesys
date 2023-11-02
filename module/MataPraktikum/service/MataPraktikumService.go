package service

import (
	"net/url"

	"github.com/rizghz/genesys/module/MataPraktikum/transfer"
)

type MataPraktikumService interface {
	GetSemuaMataPraktikum(query url.Values) []transfer.Response
	GetMataPraktikumSpesifik(kode string) *transfer.Response
	TambahMataPraktikum(data transfer.RequestBody) *transfer.Response
	EditMataPraktikum(kode string, data transfer.RequestBody) *transfer.Response
	HapusMataPraktikum(kode string) bool
}
