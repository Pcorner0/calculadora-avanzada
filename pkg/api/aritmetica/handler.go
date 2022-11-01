package aritmetica

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler interface {
	SumNumbers(ctx *gin.Context)
}

type handler struct {
	Service Service
}

func NewHandler(service Service) Handler {
	return &handler{
		Service: service,
	}
}

func (h handler) SumNumbers(ctx *gin.Context) {
	var req numbersRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	total, err := h.Service.SumNumbers(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		ctx.JSON(http.StatusInternalServerError, total)
		return
	}

	ctx.JSON(http.StatusOK, total)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
