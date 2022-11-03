package aritmetica

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler interface {
	SumSubtNumbers(ctx *gin.Context)
	DivideTwoNumbers(ctx *gin.Context)
	MultiplyNumbers(ctx *gin.Context)
	RootNumbers(ctx *gin.Context)
	PowerNumbers(ctx *gin.Context)
}

type handler struct {
	Service Service
}

func NewHandler(service Service) Handler {
	return &handler{
		Service: service,
	}
}

func (h handler) SumSubtNumbers(ctx *gin.Context) {
	var req numbersRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	total, err := h.Service.SumSubtNumbers(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, total)
}

func (h handler) DivideTwoNumbers(ctx *gin.Context) {
	var req divideRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	total, err := h.Service.DivideTwoNumbers(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, total)
}

func (h handler) MultiplyNumbers(ctx *gin.Context) {
	var req multiplyRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	total, err := h.Service.MultiplyNumbers(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, total)
}

func (h handler) RootNumbers(ctx *gin.Context) {
	var req rootRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	total, err := h.Service.RootNumbers(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, total)
}

func (h handler) PowerNumbers(ctx *gin.Context) {
	var req powerRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	total, err := h.Service.PowerNumbers(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, total)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
