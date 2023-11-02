package service

import (
	"net/url"

	"github.com/rizghz/genesys/module/Asisten/transfer"
)

type AsistenService interface {
	GetSemuaAsisten(query url.Values) []transfer.Response
	GetAsistenSpesifik(nias string) *transfer.Response
	TambahAsisten(data transfer.RequestBody) *transfer.Response
	EditAsisten(nias string, data transfer.RequestBody) *transfer.Response
	HapusAsisten(nias string) bool
}
