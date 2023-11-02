package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	jwt "github.com/rizghz/genesys/infrastructure/middleware/JWT"
	"github.com/rizghz/genesys/internal/helpers"
	"github.com/rizghz/genesys/module/Auth/service"
	"github.com/rizghz/genesys/module/Auth/transfer"
)

type AuthHttpHandler struct {
	srv service.AuthService
}

func NewAuthHttpHandler(srv service.AuthService) AuthHandler {
	return &AuthHttpHandler{
		srv: srv,
	}
}

func (h *AuthHttpHandler) Register() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		request := transfer.RegisterRequestBody{}
		if err := ctx.Bind(&request); err != nil {
			response := helpers.ApiResponse[any]{
				Status:  http.StatusBadRequest,
				Message: "invalid user register payload",
			}
			return ctx.JSON(http.StatusBadRequest, response)
		}
		result := h.srv.UserRegister(&request)
		if result != nil {
			response := helpers.ApiResponse[any]{
				Status:  http.StatusCreated,
				Message: "success",
				Data:    result,
			}
			return ctx.JSON(http.StatusCreated, response)
		}
		return ctx.JSON(http.StatusInternalServerError, nil)
	}
}

func (h *AuthHttpHandler) Login() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		request := transfer.LoginRequestBody{}
		if err := ctx.Bind(&request); err != nil {
			response := helpers.ApiResponse[any]{
				Status:  http.StatusBadRequest,
				Message: "invalid user login payload",
			}
			return ctx.JSON(http.StatusBadRequest, response)
		}
		result := h.srv.UserLogin(&request)
		if result != nil {
			response := helpers.ApiResponse[any]{
				Status:  http.StatusCreated,
				Message: "success",
				Data:    result,
			}
			return ctx.JSON(http.StatusCreated, response)
		}
		return ctx.JSON(http.StatusInternalServerError, nil)
	}
}

func (h *AuthHttpHandler) Refresh() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		// var response helpers.ApiResponse[any]
		// var request struct {
		// 	Token string `json:"token"`
		// }
		// ctx.Bind(request)
		// key := jwt.NewJwtKey()
		// if check, err := helpers.JwtValidate(request.Token, key.RefreshKey); !check {
		// 	response.Status = http.StatusInternalServerError
		// 	response.Message = err.Error()
		// 	return ctx.JSON(http.StatusInternalServerError, err.Error())
		// }
		// h.srv.RefreshToken(token)
		return nil
	}
}

func (h *AuthHttpHandler) Logout() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		token := helpers.GetJwtToken(ctx)[1]
		claim := helpers.JwtDecode[jwt.JwtPayload](&token)
		id, err := strconv.Atoi(claim.User)
		if h.srv.UserLogout(id) && err != nil {
			response := helpers.ApiResponse[any]{
				Status:  http.StatusOK,
				Message: "success",
			}
			return ctx.JSON(http.StatusOK, response)
		}
		return ctx.JSON(http.StatusInternalServerError, nil)
	}
}
