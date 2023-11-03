package repository

import (
	"net/url"

	praktikum "github.com/rizghz/genesys/module/Praktikum"
)

type PraktikumEntity praktikum.PraktikumEntity
type PraktikumModel praktikum.PraktikumModel

type PraktikumRepository interface {
	Get(query url.Values) []PraktikumEntity
	Find(id string) *PraktikumEntity
	Create(data *PraktikumModel) *PraktikumEntity
	Update(id string, data *PraktikumModel) *PraktikumEntity
	Delete(id string) bool
}
