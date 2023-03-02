package repozitory

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/mauryasaurav/go-authentication/domain/entity"
	"github.com/mauryasaurav/go-authentication/domain/interfaces"
	"github.com/mauryasaurav/go-authentication/utils/helpers"
	"gorm.io/gorm"
)

type questionRepository struct {
	db *gorm.DB
}

func NewQuestionRepository(db *gorm.DB) interfaces.QuestionRepository {
	return &questionRepository{db: db}
}

func (r *questionRepository) Create(ctx *gin.Context, question entity.QuestionSchema) (*entity.QuestionSchema, error) {
	userId, role := helpers.GetUserDeatil(ctx)
	question.Role = role
	question.UserID = userId
	result := r.db.WithContext(context.Background()).Create(&question)
	return &question, result.Error
}

func (r *questionRepository) List(ctx *gin.Context) ([]*entity.QuestionSchema, error) {
	userId, role := helpers.GetUserDeatil(ctx)
	var questions []*entity.QuestionSchema
	result := r.db.WithContext(context.Background()).Model(&entity.QuestionSchema{}).Where("user_id = ? and role = ?", userId, role).Find(&questions)
	return questions, result.Error
}
