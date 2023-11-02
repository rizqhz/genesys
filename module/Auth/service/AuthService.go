package service

import "github.com/rizghz/genesys/module/Auth/transfer"

type AuthService interface {
	UserRegister(request *transfer.RegisterRequestBody) *transfer.RegisterResponse
	UserLogin(request *transfer.LoginRequestBody) *transfer.LoginResponse
	UserLogout(id int) bool
}
