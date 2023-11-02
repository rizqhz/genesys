package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rizghz/genesys/internal/helpers"
	"github.com/rizghz/genesys/module/MataPraktikum/service"
	"github.com/rizghz/genesys/module/MataPraktikum/transfer"
)

type MataPraktikumHttpHandler struct {
	srv service.MataPraktikumService
}

func NewMataPraktikumHttpHandler(srv service.MataPraktikumService) MataPraktikumHandler {
	return &MataPraktikumHttpHandler{
		srv: srv,
	}
}

func (h *MataPraktikumHttpHandler) Index() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		query := ctx.QueryParams()
		result := h.srv.GetSemuaMataPraktikum(query)
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

func (h *MataPraktikumHttpHandler) Observe() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		kode := ctx.Param("kode")
		result := h.srv.GetMataPraktikumSpesifik(kode)
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

func (h *MataPraktikumHttpHandler) Store() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		request := transfer.RequestBody{}
		if err := ctx.Bind(&request); err != nil {
			response := helpers.ApiResponse[any]{
				Status:  http.StatusBadRequest,
				Message: "invalid mata praktikum data payload",
			}
			return ctx.JSON(http.StatusBadRequest, response)
		}
		result := h.srv.TambahMataPraktikum(request)
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

func (h *MataPraktikumHttpHandler) Edit() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		kode := ctx.Param("kode")
		request := transfer.RequestBody{}
		if err := ctx.Bind(&request); err != nil {
			response := helpers.ApiResponse[any]{
				Status:  http.StatusBadRequest,
				Message: "invalid mata praktikum data payload",
			}
			return ctx.JSON(http.StatusBadRequest, response)
		}
		result := h.srv.EditMataPraktikum(kode, request)
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

func (h *MataPraktikumHttpHandler) Destroy() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		kode := ctx.Param("kode")
		if h.srv.HapusMataPraktikum(kode) {
			response := helpers.ApiResponse[any]{
				Status:  http.StatusNoContent,
				Message: "success",
			}
			return ctx.JSON(http.StatusNoContent, response)
		}
		return ctx.JSON(http.StatusInternalServerError, nil)
	}
}
