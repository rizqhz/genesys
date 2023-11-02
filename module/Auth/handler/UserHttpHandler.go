package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/rizghz/genesys/internal/helpers"
	"github.com/rizghz/genesys/module/Auth/service"
	"github.com/rizghz/genesys/module/Auth/transfer"
	"github.com/sirupsen/logrus"
)

type UserHttpHandler struct {
	srv service.UserService
}

func NewUserHttpHandler(srv service.UserService) UserHandler {
	return &UserHttpHandler{
		srv: srv,
	}
}

func (h *UserHttpHandler) Index() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		result := h.srv.GetSemuaUser(ctx)
		if len(result) != 0 {
			response := helpers.ApiResponse[any]{
				Status:  http.StatusOK,
				Message: "success",
				Data:    result,
			}
			return ctx.JSON(http.StatusOK, response)
		}
		return ctx.JSON(http.StatusNoContent, nil)
	}
}

func (h *UserHttpHandler) Observe() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			response := helpers.ApiResponse[any]{
				Status:  http.StatusBadRequest,
				Message: "invalid user id",
			}
			return ctx.JSON(http.StatusBadRequest, response)
		}
		result := h.srv.GetUserSpesifik(id)
		if result != nil {
			response := helpers.ApiResponse[any]{
				Status:  http.StatusOK,
				Message: "success",
				Data:    result,
			}
			return ctx.JSON(http.StatusOK, response)
		}
		return ctx.JSON(http.StatusNoContent, nil)
	}
}

func (h *UserHttpHandler) Store() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		request := &transfer.UserRequestBody{}
		if err := ctx.Bind(request); err != nil {
			response := helpers.ApiResponse[any]{
				Status:  http.StatusBadRequest,
				Message: "invalid user data payload",
			}
			return ctx.JSON(http.StatusBadRequest, response)
		}
		file, err := ctx.FormFile("foto")
		if err != nil {
			logrus.Error(err.Error())
		} else {
			url := h.srv.UploadFoto(file)
			if url != nil {
				request.Foto = *url
			}
		}
		result := h.srv.TambahUser(request)
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

func (h *UserHttpHandler) Edit() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			response := helpers.ApiResponse[any]{
				Status:  http.StatusBadRequest,
				Message: "invalid user id",
			}
			return ctx.JSON(http.StatusBadRequest, response)
		}
		request := &transfer.UserRequestBody{}
		if err := ctx.Bind(request); err != nil {
			response := helpers.ApiResponse[any]{
				Status:  http.StatusBadRequest,
				Message: "invalid user data payload",
			}
			return ctx.JSON(http.StatusBadRequest, response)
		}
		file, err := ctx.FormFile("foto")
		if err != nil {
			logrus.Error(err.Error())
		} else {
			url := h.srv.UploadFoto(file)
			if url != nil {
				request.Foto = *url
			}
		}
		result := h.srv.EditUser(id, request)
		if result != nil {
			response := helpers.ApiResponse[any]{
				Status:  http.StatusOK,
				Message: "success",
				Data:    result,
			}
			return ctx.JSON(http.StatusOK, response)
		}
		return ctx.JSON(http.StatusInternalServerError, nil)
	}
}

func (h *UserHttpHandler) Destroy() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			response := helpers.ApiResponse[any]{
				Status:  http.StatusBadRequest,
				Message: "invalid user id",
			}
			return ctx.JSON(http.StatusBadRequest, response)
		}
		if h.srv.HapusUser(id) {
			response := helpers.ApiResponse[any]{
				Status:  http.StatusNoContent,
				Message: "success",
			}
			return ctx.JSON(http.StatusNoContent, response)
		}
		return ctx.JSON(http.StatusInternalServerError, nil)
	}
}
