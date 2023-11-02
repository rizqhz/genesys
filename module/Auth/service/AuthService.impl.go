package service

import (
	"fmt"

	jwt "github.com/rizghz/genesys/infrastructure/middleware/JWT"
	auth "github.com/rizghz/genesys/module/Auth"
	"github.com/rizghz/genesys/module/Auth/repository"
	"github.com/rizghz/genesys/module/Auth/transfer"
)

type AuthServiceImpl struct {
	repo repository.CredentialRepository
}

func NewAuthServiceImpl(r repository.CredentialRepository) AuthService {
	return &AuthServiceImpl{
		repo: r,
	}
}

func (srv *AuthServiceImpl) UserRegister(request *transfer.RegisterRequestBody) *transfer.RegisterResponse {
	data := &auth.UserModel{
		Nama:    request.Nama,
		Email:   request.Email,
		Telepon: request.Telepon,
		Credential: auth.CredentialModel{
			Usercode: request.Usercode,
			Password: request.Password,
			Role:     request.Role,
		},
	}
	return srv.repo.Create(data)
}

func (srv *AuthServiceImpl) UserLogin(request *transfer.LoginRequestBody) *transfer.LoginResponse {
	param := &auth.CredentialModel{
		Usercode: request.Usercode,
		Password: request.Password,
	}
	result := srv.repo.Search(param)
	if result != nil {
		token := jwt.NewJwtToken(fmt.Sprint(result.ID), result.Credential.Role)
		token.Generate()
		result.Credential.Token = token.RefreshToken
		if srv.repo.Update(result) == nil {
			return nil
		}
		return &transfer.LoginResponse{
			AccessToken:  token.AccessToken,
			RefreshToken: token.RefreshToken,
		}
	}
	return nil
}

func (srv *AuthServiceImpl) UserLogout(id int) bool {
	data := &auth.CredentialModel{
		ID: id,
	}
	return srv.repo.Delete(data)
}
