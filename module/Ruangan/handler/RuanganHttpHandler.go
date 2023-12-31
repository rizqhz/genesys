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
		result := h.srv.GetSemuaRuangan(ctx)
		if ctx.Get("authorization.error") != nil {
			response := helpers.ApiResponse[any]{
				Status:  http.StatusUnauthorized,
				Message: "user bukan admin",
			}
			return ctx.JSON(http.StatusUnauthorized, response)
		}
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
		result := h.srv.GetRuanganSpesifik(ctx, kode)
		if ctx.Get("authorization.error") != nil {
			response := helpers.ApiResponse[any]{
				Status:  http.StatusUnauthorized,
				Message: "user bukan admin",
			}
			return ctx.JSON(http.StatusUnauthorized, response)
		}
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
		request := &transfer.RequestBody{}
		if err := ctx.Bind(request); err != nil {
			response := helpers.ApiResponse[any]{
				Status:  http.StatusBadRequest,
				Message: "invalid ruangan data payload",
			}
			return ctx.JSON(http.StatusBadRequest, response)
		}
		result := h.srv.TambahRuangan(ctx, request)
		if ctx.Get("authorization.error") != nil {
			response := helpers.ApiResponse[any]{
				Status:  http.StatusUnauthorized,
				Message: "user bukan admin",
			}
			return ctx.JSON(http.StatusUnauthorized, response)
		}
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
		request := &transfer.RequestBody{}
		if err := ctx.Bind(request); err != nil {
			response := helpers.ApiResponse[any]{
				Status:  http.StatusBadRequest,
				Message: "invalid ruangan data payload",
			}
			return ctx.JSON(http.StatusBadRequest, response)
		}
		result := h.srv.EditRuangan(ctx, kode, request)
		if ctx.Get("authorization.error") != nil {
			response := helpers.ApiResponse[any]{
				Status:  http.StatusUnauthorized,
				Message: "user bukan admin",
			}
			return ctx.JSON(http.StatusUnauthorized, response)
		}
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
		if h.srv.HapusRuangan(ctx, kode) {
			response := helpers.ApiResponse[any]{
				Status:  http.StatusNoContent,
				Message: "success",
			}
			return ctx.JSON(http.StatusNoContent, response)
		}
		if ctx.Get("authorization.error") != nil {
			response := helpers.ApiResponse[any]{
				Status:  http.StatusUnauthorized,
				Message: "user bukan admin",
			}
			return ctx.JSON(http.StatusUnauthorized, response)
		}
		return ctx.JSON(http.StatusInternalServerError, nil)
	}
}
