package repozitory

import (
	"errors"

	"github.com/mauryasaurav/go-authentication/domain/entity"
	"github.com/mauryasaurav/go-authentication/domain/interfaces"
	"gorm.io/gorm"
)

const dbTableName string = "users"

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) interfaces.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUsers(user entity.UserSchema) (*entity.UserSchema, error) {
	err := r.db.Table(dbTableName).
		Create(&user).
		Error
	if err != nil {
		return nil, err
	}
	return &user, err
}

func (r *userRepository) FindByEmail(email string) (*entity.UserSchema, error) {
	user := new(entity.UserSchema)
	err := r.db.Table(dbTableName).
		Where("email = ? ", email).
		First(&user).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) UpdateByEmail(email string, user entity.UserSchema) error {
	err := r.db.Table(dbTableName).
		Where("email = ? ", email).
		Updates(user).
		Error
	return err
}
