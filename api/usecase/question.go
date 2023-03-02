package usecase

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mauryasaurav/go-authentication/domain/dto"
	"github.com/mauryasaurav/go-authentication/domain/entity"
	"github.com/mauryasaurav/go-authentication/domain/interfaces"
)

type questionUsecase struct {
	questionRepo interfaces.QuestionRepository
}

func NewQuestionUsecase(questionRepo interfaces.QuestionRepository) interfaces.QuestionUsecase {
	return &questionUsecase{
		questionRepo: questionRepo,
	}
}

func (u *questionUsecase) Create(ctx *gin.Context, req *dto.QuestionValidator) error {
	if _, err := u.questionRepo.Create(ctx, entity.QuestionSchema{
		Question:   req.Question,
		Answer:     req.Answer,
		CategoryID: uuid.MustParse(req.CategoryID),
	}); err != nil {
		return err
	}

	return nil
}

func (u *questionUsecase) List(ctx *gin.Context) ([]*entity.QuestionSchema, error) {
	list, err := u.questionRepo.List(ctx)
	if err != nil {
		return nil, err
	}

	return list, nil
}
