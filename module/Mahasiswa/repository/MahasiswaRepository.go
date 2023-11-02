package repository

import (
	"net/url"

	mahasiswa "github.com/rizghz/genesys/module/Mahasiswa"
)

type Entity mahasiswa.MahasiswaEntity
type Model mahasiswa.MahasiswaModel

type MahasiswaRepository interface {
	Get(query url.Values) []Entity
	Find(npm string) *Entity
	Create(data *Model) *Entity
	Update(npm string, data *Model) *Entity
	Delete(npm string) bool
}
