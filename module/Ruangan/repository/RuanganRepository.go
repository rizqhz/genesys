package repository

import (
	"net/url"

	ruangan "github.com/rizghz/genesys/module/Ruangan"
)

type Entity ruangan.RuanganEntity
type Model ruangan.RuanganModel

type RuanganRepository interface {
	Get(query url.Values) []Entity
	Find(kode string) *Entity
	Create(data *Model) *Entity
	Update(kode string, data *Model) *Entity
	Delete(kode string) bool
}
