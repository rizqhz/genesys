package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rizghz/genesys/internal/helpers"
	"github.com/rizghz/genesys/module/Asisten/service"
	"github.com/rizghz/genesys/module/Asisten/transfer"
)

type AsistenHttpHandler struct {
	srv service.AsistenService
}

func NewAsistenHttpHandler(srv service.AsistenService) AsistenHandler {
	return &AsistenHttpHandler{
		srv: srv,
	}
}

func (h *AsistenHttpHandler) Index() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		query := ctx.QueryParams()
		result := h.srv.GetSemuaAsisten(query)
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

func (h *AsistenHttpHandler) Observe() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		kode := ctx.Param("kode")
		result := h.srv.GetAsistenSpesifik(kode)
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

func (h *AsistenHttpHandler) Store() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		request := transfer.RequestBody{}
		if err := ctx.Bind(&request); err != nil {
			response := helpers.ApiResponse[any]{
				Status:  http.StatusBadRequest,
				Message: "invalid asisten data payload",
			}
			return ctx.JSON(http.StatusBadRequest, response)
		}
		result := h.srv.TambahAsisten(request)
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

func (h *AsistenHttpHandler) Edit() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		kode := ctx.Param("kode")
		request := transfer.RequestBody{}
		if err := ctx.Bind(&request); err != nil {
			response := helpers.ApiResponse[any]{
				Status:  http.StatusBadRequest,
				Message: "invalid asisten data payload",
			}
			return ctx.JSON(http.StatusBadRequest, response)
		}
		result := h.srv.EditAsisten(kode, request)
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

func (h *AsistenHttpHandler) Destroy() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		kode := ctx.Param("kode")
		if h.srv.HapusAsisten(kode) {
			response := helpers.ApiResponse[any]{
				Status:  http.StatusNoContent,
				Message: "success",
			}
			return ctx.JSON(http.StatusNoContent, response)
		}
		return ctx.JSON(http.StatusInternalServerError, nil)
	}
}
