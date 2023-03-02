package repozitory

import (
	"context"

	"github.com/mauryasaurav/go-authentication/domain/entity"
	"github.com/mauryasaurav/go-authentication/domain/interfaces"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) interfaces.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUsers(user entity.UserSchema) (*entity.UserSchema, error) {
	result := r.db.WithContext(context.Background()).Create(&user)
	return &user, result.Error
}

func (r *userRepository) FindByEmail(email string) (*entity.UserSchema, error) {
	var user entity.UserSchema
	result := r.db.WithContext(context.Background()).Model(&entity.UserSchema{}).Where("email = ?", email).Find(&user).Limit(1)
	return &user, result.Error
}

func (r *userRepository) UpdateByEmail(email string, user entity.UserSchema) error {
	result := r.db.WithContext(context.Background()).Where("email = ?", email).Updates(user)
	result.Scan(&user)
	return result.Error
}
