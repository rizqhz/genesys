package repository

import (
	"net/url"

	asisten "github.com/rizghz/genesys/module/Asisten"
)

type AsistenEntity asisten.AsistenEntity
type AsistenModel asisten.AsistenModel

type AsistenRepository interface {
	Get(query url.Values) []AsistenEntity
	Find(kode string) *AsistenEntity
	Create(data *AsistenModel) *AsistenEntity
	Update(kode string, data *AsistenModel) *AsistenEntity
	Delete(kode string) bool
}
