package repository

import (
	"net/url"

	matkum "github.com/rizghz/genesys/module/MataPraktikum"
)

type MatkumEntity matkum.MataPraktikum
type MatkumModel matkum.MataPraktikumModel

type MatkumRepository interface {
	Get(query url.Values) []MatkumEntity
	Find(kode string) *MatkumEntity
	Create(data *MatkumModel) *MatkumEntity
	Update(kode string, data *MatkumModel) *MatkumEntity
	Delete(kode string) bool
}
