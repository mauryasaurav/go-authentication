package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mauryasaurav/go-authentication/domain/dto"
	"github.com/mauryasaurav/go-authentication/domain/interfaces"
)

type categoryHandler struct {
	categoryUsecase interfaces.CategoryUsecase
}

func NewCategoryHandler(route *gin.RouterGroup, u interfaces.CategoryUsecase) {
	handler := categoryHandler{categoryUsecase: u}
	route.POST("/category", handler.Create)
	route.GET("/categories", handler.List)
}

func (h *categoryHandler) Create(ctx *gin.Context) {
	created := new(dto.CategoryValidator)
	if err := ctx.Bind(created); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.categoryUsecase.Create(ctx, created)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "category created successfully"})
}

func (h *categoryHandler) List(ctx *gin.Context) {
	list, err := h.categoryUsecase.List(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": list})
}
