package usecase

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/mauryasaurav/go-authentication/domain/entity"
	"github.com/mauryasaurav/go-authentication/domain/interfaces"
)

type publicUsecase struct {
	pubicRepo interfaces.PublicRepository
}

func NewPublicUsecase(pubicRepo interfaces.PublicRepository) interfaces.PublicUsecase {
	return &publicUsecase{
		pubicRepo: pubicRepo,
	}
}

func (u *publicUsecase) ListQuestions(ctx *gin.Context) ([]*entity.QuestionSchema, error) {
	questions, err := u.pubicRepo.ListQuestions(context.Background())
	if err != nil {
		return nil, err
	}

	return questions, nil
}

func (u *publicUsecase) ListCategories(ctx *gin.Context) ([]*entity.CategorySchema, error) {
	categories, err := u.pubicRepo.ListCategories(context.Background())

	if err != nil {
		return nil, err
	}

	return categories, nil
}
