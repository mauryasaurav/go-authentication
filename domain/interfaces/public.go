package interfaces

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/mauryasaurav/go-authentication/domain/entity"
)

type PublicUsecase interface {
	ListCategories(ctx *gin.Context) ([]*entity.CategorySchema, error)
	ListQuestions(ctx *gin.Context) ([]*entity.QuestionSchema, error)
}

type PublicRepository interface {
	ListCategories(ctx context.Context) ([]*entity.CategorySchema, error)
	ListQuestions(ctx context.Context) ([]*entity.QuestionSchema, error)
}
