package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rizghz/genesys/internal/helpers"
	"github.com/rizghz/genesys/module/MataPraktikum/service"
	"github.com/rizghz/genesys/module/MataPraktikum/transfer"
)

type MatkumHttpHandler struct {
	srv service.MatkumService
}

func NewMatkumHttpHandler(srv service.MatkumService) MatkumHandler {
	return &MatkumHttpHandler{
		srv: srv,
	}
}

func (h *MatkumHttpHandler) Index() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		result := h.srv.GetSemuaMatkum(ctx)
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

func (h *MatkumHttpHandler) Observe() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		kode := ctx.Param("kode")
		result := h.srv.GetMatkumSpesifik(ctx, kode)
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

func (h *MatkumHttpHandler) Store() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		request := &transfer.RequestBody{}
		if err := ctx.Bind(request); err != nil {
			response := helpers.ApiResponse[any]{
				Status:  http.StatusBadRequest,
				Message: "invalid mata praktikum data payload",
			}
			return ctx.JSON(http.StatusBadRequest, response)
		}
		result := h.srv.TambahMatkum(ctx, request)
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

func (h *MatkumHttpHandler) Edit() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		kode := ctx.Param("kode")
		request := &transfer.RequestBody{}
		if err := ctx.Bind(request); err != nil {
			response := helpers.ApiResponse[any]{
				Status:  http.StatusBadRequest,
				Message: "invalid mata praktikum data payload",
			}
			return ctx.JSON(http.StatusBadRequest, response)
		}
		result := h.srv.EditMatkum(ctx, kode, request)
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

func (h *MatkumHttpHandler) Destroy() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		kode := ctx.Param("kode")
		if h.srv.HapusMatkum(ctx, kode) {
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
