package repository

import (
	"net/url"

	auth "github.com/rizghz/genesys/module/Auth"
)

type UserEntity auth.UserEntity
type UserModel auth.UserModel

type UserRepository interface {
	Get(query url.Values) []UserEntity
	Find(id int) *UserEntity
	Create(data *UserModel) *UserEntity
	Update(data *UserModel) *UserEntity
	Delete(id int) bool
}
