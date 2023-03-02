package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mauryasaurav/go-authentication/domain/dto"
	"github.com/mauryasaurav/go-authentication/domain/interfaces"
)

type questionHandler struct {
	questionUsecase interfaces.QuestionUsecase
}

func NewQuestionHandler(route *gin.RouterGroup, u interfaces.QuestionUsecase) {
	handler := questionHandler{questionUsecase: u}
	route.POST("/question", handler.Create)
	route.GET("/questions", handler.List)
}

func (h *questionHandler) Create(ctx *gin.Context) {
	created := new(dto.QuestionValidator)
	if err := ctx.Bind(created); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.questionUsecase.Create(ctx, created)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "question created successfully"})
}

func (h *questionHandler) List(ctx *gin.Context) {
	list, err := h.questionUsecase.List(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": list})
}
