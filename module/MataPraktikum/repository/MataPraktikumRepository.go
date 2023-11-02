package repository

import (
	"net/url"

	matkum "github.com/rizghz/genesys/module/MataPraktikum"
)

type (
	Entity matkum.MataPraktikum
	Model  matkum.MataPraktikumModel
)

type MataPraktikumRepository interface {
	Get(query url.Values) []Entity
	Find(kode string) *Entity
	Create(data *Model) *Entity
	Update(kode string, data *Model) *Entity
	Delete(kode string) bool
}
