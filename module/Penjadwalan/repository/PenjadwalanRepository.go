package repository

import (
	"net/url"

	penjadwalan "github.com/rizghz/genesys/module/Penjadwalan"
)

type JadwalEntity penjadwalan.PenjadwalanEntity
type JadwalModel penjadwalan.PenjadwalanModel

type PenjadwalanRepository interface {
	Get(query url.Values) []JadwalEntity
	Find(id int) *JadwalEntity
	Create(data *JadwalModel) *JadwalEntity
	Update(data *JadwalModel) *JadwalEntity
	Delete(id int) bool
}
