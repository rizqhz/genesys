package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rizghz/genesys/internal/helpers"
	"github.com/rizghz/genesys/module/Mahasiswa/service"
	"github.com/rizghz/genesys/module/Mahasiswa/transfer"
)

type MahasiswaHttpHandler struct {
	srv service.MahasiswaService
}

func NewMahasiswaHttpHandler(srv service.MahasiswaService) MahasiswaHandler {
	return &MahasiswaHttpHandler{
		srv: srv,
	}
}

func (h *MahasiswaHttpHandler) Index() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		query := ctx.QueryParams()
		result := h.srv.GetSemuaMahasiswa(query)
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

func (h *MahasiswaHttpHandler) Observe() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		npm := ctx.Param("npm")
		result := h.srv.GetMahasiswaSpesifik(npm)
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

func (h *MahasiswaHttpHandler) Store() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		request := transfer.RequestBody{}
		if err := ctx.Bind(&request); err != nil {
			response := helpers.ApiResponse[any]{
				Status:  http.StatusBadRequest,
				Message: "invalid mahasiswa data payload",
			}
			return ctx.JSON(http.StatusBadRequest, response)
		}
		result := h.srv.TambahMahasiswa(request)
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

func (h *MahasiswaHttpHandler) Edit() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		npm := ctx.Param("npm")
		request := transfer.RequestBody{}
		if err := ctx.Bind(&request); err != nil {
			response := helpers.ApiResponse[any]{
				Status:  http.StatusBadRequest,
				Message: "invalid mahasiswa data payload",
			}
			return ctx.JSON(http.StatusBadRequest, response)
		}
		result := h.srv.EditMahasiswa(npm, request)
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

func (h *MahasiswaHttpHandler) Destroy() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		npm := ctx.Param("npm")
		if h.srv.HapusMahasiswa(npm) {
			response := helpers.ApiResponse[any]{
				Status:  http.StatusNoContent,
				Message: "success",
			}
			return ctx.JSON(http.StatusNoContent, response)
		}
		return ctx.JSON(http.StatusInternalServerError, nil)
	}
}
