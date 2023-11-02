package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rizghz/genesys/internal/helpers"
	"github.com/rizghz/genesys/module/Ruangan/service"
	"github.com/rizghz/genesys/module/Ruangan/transfer"
)

type RuanganHttpHandler struct {
	srv service.RuanganService
}

func NewRuanganHttpHandler(srv service.RuanganService) RuanganHandler {
	return &RuanganHttpHandler{
		srv: srv,
	}
}

func (h *RuanganHttpHandler) Index() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		query := ctx.QueryParams()
		result := h.srv.GetSemuaRuangan(query)
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

func (h *RuanganHttpHandler) Observe() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		kode := ctx.Param("kode")
		result := h.srv.GetRuanganSpesifik(kode)
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

func (h *RuanganHttpHandler) Store() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		request := transfer.RequestBody{}
		if err := ctx.Bind(&request); err != nil {
			response := helpers.ApiResponse[any]{
				Status:  http.StatusBadRequest,
				Message: "invalid ruangan data payload",
			}
			return ctx.JSON(http.StatusBadRequest, response)
		}
		result := h.srv.TambahRuangan(request)
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

func (h *RuanganHttpHandler) Edit() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		kode := ctx.Param("kode")
		request := transfer.RequestBody{}
		if err := ctx.Bind(&request); err != nil {
			response := helpers.ApiResponse[any]{
				Status:  http.StatusBadRequest,
				Message: "invalid ruangan data payload",
			}
			return ctx.JSON(http.StatusBadRequest, response)
		}
		result := h.srv.EditRuangan(kode, request)
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

func (h *RuanganHttpHandler) Destroy() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		kode := ctx.Param("kode")
		if h.srv.HapusRuangan(kode) {
			response := helpers.ApiResponse[any]{
				Status:  http.StatusNoContent,
				Message: "success",
			}
			return ctx.JSON(http.StatusNoContent, response)
		}
		return ctx.JSON(http.StatusInternalServerError, nil)
	}
}
