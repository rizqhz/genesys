package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rizghz/genesys/internal/helpers"
	"github.com/rizghz/genesys/module/Kelas/service"
	"github.com/rizghz/genesys/module/Kelas/transfer"
)

type KelasHttpHandler struct {
	srv service.KelasService
}

func NewKelasHttpHandler(srv service.KelasService) KelasHandler {
	return &KelasHttpHandler{
		srv: srv,
	}
}

func (h *KelasHttpHandler) Index() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		result := h.srv.GetSemuaKelas(ctx)
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

func (h *KelasHttpHandler) Observe() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		kode := ctx.Param("kode")
		result := h.srv.GetKelasSpesifik(ctx, kode)
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

func (h *KelasHttpHandler) Store() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		request := &transfer.RequestBody{}
		if err := ctx.Bind(request); err != nil {
			response := helpers.ApiResponse[any]{
				Status:  http.StatusBadRequest,
				Message: "invalid kelas data payload",
			}
			return ctx.JSON(http.StatusBadRequest, response)
		}
		result := h.srv.TambahKelas(ctx, request)
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

func (h *KelasHttpHandler) Edit() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		kode := ctx.Param("kode")
		request := &transfer.RequestBody{}
		if err := ctx.Bind(request); err != nil {
			response := helpers.ApiResponse[any]{
				Status:  http.StatusBadRequest,
				Message: "invalid kelas data payload",
			}
			return ctx.JSON(http.StatusBadRequest, response)
		}
		result := h.srv.EditKelas(ctx, kode, request)
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

func (h *KelasHttpHandler) Destroy() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		kode := ctx.Param("kode")
		if h.srv.HapusKelas(ctx, kode) {
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
