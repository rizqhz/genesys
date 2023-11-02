package service

import (
	"net/url"

	"github.com/rizghz/genesys/module/Mahasiswa/transfer"
)

type MahasiswaService interface {
	GetSemuaMahasiswa(query url.Values) []transfer.Response
	GetMahasiswaSpesifik(npm string) *transfer.Response
	TambahMahasiswa(data transfer.RequestBody) *transfer.Response
	EditMahasiswa(npm string, data transfer.RequestBody) *transfer.Response
	HapusMahasiswa(npm string) bool
}
