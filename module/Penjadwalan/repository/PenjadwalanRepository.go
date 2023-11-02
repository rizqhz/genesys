package repository

import (
	"net/url"

	penjadwalan "github.com/rizghz/genesys/module/Penjadwalan"
)

type Entity penjadwalan.PenjadwalanEntity
type Model penjadwalan.PenjadwalanModel

type PenjadwalanRepository interface {
	Get(query url.Values) []Entity
	Find(id int) *Entity
	Create(data *Model) *Entity
	Update(data *Model) *Entity
	Delete(id int) bool
}
