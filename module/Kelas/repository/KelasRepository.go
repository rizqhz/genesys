package repository

import (
	"net/url"

	kelas "github.com/rizghz/genesys/module/Kelas"
)

type Entity kelas.KelasEntity
type Model kelas.KelasModel

type KelasRepository interface {
	Get(query url.Values) []Entity
	Find(kode string) *Entity
	Create(data *Model) *Entity
	Update(kode string, data *Model) *Entity
	Delete(kode string) bool
}
