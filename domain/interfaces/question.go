package interfaces

import (
	"github.com/gin-gonic/gin"
	"github.com/mauryasaurav/go-authentication/domain/dto"
	"github.com/mauryasaurav/go-authentication/domain/entity"
)

type QuestionUsecase interface {
	Create(ctx *gin.Context, request *dto.QuestionValidator) error
	List(ctx *gin.Context) ([]*entity.QuestionSchema, error)
}

type QuestionRepository interface {
	Create(ctx *gin.Context, user entity.QuestionSchema) (*entity.QuestionSchema, error)
	List(ctx *gin.Context) ([]*entity.QuestionSchema, error)
}
