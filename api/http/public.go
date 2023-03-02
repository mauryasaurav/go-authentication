package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mauryasaurav/go-authentication/domain/interfaces"
)

type publicHandler struct {
	publicUsecase interfaces.PublicUsecase
}

func NewPublicHandler(route *gin.RouterGroup, u interfaces.PublicUsecase) {
	handler := publicHandler{publicUsecase: u}
	route.GET("/questions", handler.ListQuestions)
	route.GET("/categories", handler.ListCategories)
}

func (h *publicHandler) ListQuestions(ctx *gin.Context) {
	list, err := h.publicUsecase.ListQuestions(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": list})
}

func (h *publicHandler) ListCategories(ctx *gin.Context) {
	list, err := h.publicUsecase.ListCategories(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": list})
}
