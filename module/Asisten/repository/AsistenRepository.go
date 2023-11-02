package repository

import (
	"net/url"

	asisten "github.com/rizghz/genesys/module/Asisten"
)

type Entity asisten.AsistenEntity
type Model asisten.AsistenModel

type AsistenRepository interface {
	Get(query url.Values) []Entity
	Find(kode string) *Entity
	Create(data *Model) *Entity
	Update(kode string, data *Model) *Entity
	Delete(kode string) bool
}
