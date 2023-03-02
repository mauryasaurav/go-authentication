package interfaces

import (
	"github.com/gin-gonic/gin"
	"github.com/mauryasaurav/go-authentication/domain/dto"
	"github.com/mauryasaurav/go-authentication/domain/entity"
)

type CategoryUsecase interface {
	Create(ctx *gin.Context, request *dto.CategoryValidator) error
	List(ctx *gin.Context) ([]*entity.CategorySchema, error)
}

type CategoryRepository interface {
	Create(ctx *gin.Context, user entity.CategorySchema) (*entity.CategorySchema, error)
	List(ctx *gin.Context) ([]*entity.CategorySchema, error)
}
