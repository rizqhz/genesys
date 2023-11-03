package repository

import (
	"net/url"

	ruangan "github.com/rizghz/genesys/module/Ruangan"
)

type RuanganEntity ruangan.RuanganEntity
type RuanganModel ruangan.RuanganModel

type RuanganRepository interface {
	Get(query url.Values) []RuanganEntity
	Find(kode string) *RuanganEntity
	Create(data *RuanganModel) *RuanganEntity
	Update(kode string, data *RuanganModel) *RuanganEntity
	Delete(kode string) bool
}
