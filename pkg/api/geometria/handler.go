package geometria

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler interface {
	CosNumbers(ctx *gin.Context)
	SinNumbers(ctx *gin.Context)
}

type handler struct {
	Service Service
}

func NewHandler(service Service) Handler {
	return &handler{
		Service: service,
	}
}

func (h handler) SinNumbers(ctx *gin.Context) {
	var req sinRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	total, err := h.Service.SinNumbers(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, total)
}

func (h handler) CosNumbers(ctx *gin.Context) {
	var req cosRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	total, err := h.Service.CosNumbers(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, total)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}