package repozitory

import (
	"context"

	"github.com/mauryasaurav/go-authentication/domain/entity"
	"github.com/mauryasaurav/go-authentication/domain/interfaces"
	"gorm.io/gorm"
)

type publicRepository struct {
	db *gorm.DB
}

func NewPublicRepository(db *gorm.DB) interfaces.PublicRepository {
	return &publicRepository{db: db}
}

func (r *publicRepository) ListCategories(ctx context.Context) ([]*entity.CategorySchema, error) {
	var categories []*entity.CategorySchema
	result := r.db.WithContext(ctx).Model(&entity.CategorySchema{}).Where("role = 'admin'").Find(&categories)
	return categories, result.Error
}

func (r *publicRepository) ListQuestions(ctx context.Context) ([]*entity.QuestionSchema, error) {
	var questions []*entity.QuestionSchema
	result := r.db.WithContext(ctx).Model(&entity.QuestionSchema{}).Where("role = 'admin'").Find(&questions)
	return questions, result.Error
}
