package repository

import (
	"net/url"

	kelas "github.com/rizghz/genesys/module/Kelas"
)

type KelasEntity kelas.KelasEntity
type KelasModel kelas.KelasModel

type KelasRepository interface {
	Get(query url.Values) []KelasEntity
	Find(kode string) *KelasEntity
	Create(data *KelasModel) *KelasEntity
	Update(kode string, data *KelasModel) *KelasEntity
	Delete(kode string) bool
}
