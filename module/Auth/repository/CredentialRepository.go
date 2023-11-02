package repository

import (
	auth "github.com/rizghz/genesys/module/Auth"
	"github.com/rizghz/genesys/module/Auth/transfer"
)

type CredentialRepository interface {
	Create(data *auth.UserModel) *transfer.RegisterResponse
	Update(data *auth.UserModel) *auth.UserModel
	Search(data *auth.CredentialModel) *auth.UserModel
	Delete(data *auth.CredentialModel) bool
}
