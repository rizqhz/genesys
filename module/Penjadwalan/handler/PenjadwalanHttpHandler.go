package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/rizghz/genesys/internal/helpers"
	"github.com/rizghz/genesys/module/Penjadwalan/service"
	"github.com/rizghz/genesys/module/Penjadwalan/transfer"
)

type PenjadwalanHttpHandler struct {
	srv service.PenjadwalanService
}

func NewPenjadwalanHttpHandler(srv service.PenjadwalanService) PenjadwalanHandler {
	return &PenjadwalanHttpHandler{
		srv: srv,
	}
}

func (h *PenjadwalanHttpHandler) Index() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		result := h.srv.GetSemuaJadwal(ctx)
		if ctx.Get("authorization.error") != nil {
			response := helpers.ApiResponse[any]{
				Status:  http.StatusUnauthorized,
				Message: "user bukan admin jadwal",
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

func (h *PenjadwalanHttpHandler) Observe() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			response := helpers.ApiResponse[any]{
				Status:  http.StatusBadRequest,
				Message: "invalid jadwal id",
			}
			return ctx.JSON(http.StatusBadRequest, response)
		}
		result := h.srv.GetJadwalSpesifik(ctx, id)
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

func (h *PenjadwalanHttpHandler) Store() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		request := &transfer.RequestBody{}
		if err := ctx.Bind(request); err != nil {
			response := helpers.ApiResponse[any]{
				Status:  http.StatusBadRequest,
				Message: "invalid jadwal data payload",
			}
			return ctx.JSON(http.StatusBadRequest, response)
		}
		result := h.srv.TambahJadwal(ctx, request)
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

func (h *PenjadwalanHttpHandler) Edit() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			response := helpers.ApiResponse[any]{
				Status:  http.StatusBadRequest,
				Message: "invalid jadwal id",
			}
			return ctx.JSON(http.StatusBadRequest, response)
		}
		request := &transfer.RequestBody{}
		if err := ctx.Bind(request); err != nil {
			response := helpers.ApiResponse[any]{
				Status:  http.StatusBadRequest,
				Message: "invalid jadwal data payload",
			}
			return ctx.JSON(http.StatusBadRequest, response)
		}
		result := h.srv.EditJadwal(ctx, id, request)
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

func (h *PenjadwalanHttpHandler) Destroy() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			response := helpers.ApiResponse[any]{
				Status:  http.StatusBadRequest,
				Message: "invalid jadwal id",
			}
			return ctx.JSON(http.StatusBadRequest, response)
		}
		if h.srv.HapusJadwal(ctx, id) {
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
