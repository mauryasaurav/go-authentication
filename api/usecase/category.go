package usecase

import (
	"github.com/gin-gonic/gin"
	"github.com/mauryasaurav/go-authentication/domain/dto"
	"github.com/mauryasaurav/go-authentication/domain/entity"
	"github.com/mauryasaurav/go-authentication/domain/interfaces"
)

type categoryUsecase struct {
	categoryRepo interfaces.CategoryRepository
}

func NewCategoryUsecase(categoryRepo interfaces.CategoryRepository) interfaces.CategoryUsecase {
	return &categoryUsecase{
		categoryRepo: categoryRepo,
	}
}

func (u *categoryUsecase) Create(ctx *gin.Context, req *dto.CategoryValidator) error {
	if _, err := u.categoryRepo.Create(ctx, entity.CategorySchema{
		Name: req.Name,
	}); err != nil {
		return err
	}

	return nil
}

func (u *categoryUsecase) List(ctx *gin.Context) ([]*entity.CategorySchema, error) {
	list, err := u.categoryRepo.List(ctx)
	if err != nil {
		return nil, err
	}

	return list, nil
}
