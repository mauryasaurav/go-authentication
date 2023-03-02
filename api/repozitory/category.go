package repozitory

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/mauryasaurav/go-authentication/domain/entity"
	"github.com/mauryasaurav/go-authentication/domain/interfaces"
	"github.com/mauryasaurav/go-authentication/utils/helpers"
	"gorm.io/gorm"
)

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) interfaces.CategoryRepository {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) Create(ctx *gin.Context, category entity.CategorySchema) (*entity.CategorySchema, error) {
	userId, role := helpers.GetUserDeatil(ctx)
	category.Role = role
	category.UserID = userId
	result := r.db.WithContext(context.Background()).Create(&category)
	return &category, result.Error
}

func (r *categoryRepository) List(ctx *gin.Context) ([]*entity.CategorySchema, error) {
	userId, role := helpers.GetUserDeatil(ctx)
	var categories []*entity.CategorySchema
	result := r.db.WithContext(context.Background()).Model(&entity.CategorySchema{}).Where("user_id = ? and role = ?", userId, role).Find(&categories)
	return categories, result.Error
}
