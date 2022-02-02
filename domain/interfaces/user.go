package interfaces

import (
	"github.com/gin-gonic/gin"
	"github.com/mauryasaurav/go-authentication/domain/dto"
	"github.com/mauryasaurav/go-authentication/domain/entity"
)

type UserUsecase interface {
	AuthenticationUser(ctx *gin.Context, oldPassword string, currentPassword string) bool
	CreateUserHandler(ctx *gin.Context, request *dto.UserValidator) error
	UpdateUserHandler(ctx *gin.Context, request dto.UserUpdateValidator) error
	LoginUserHandler(ctx *gin.Context, request dto.UserLoginValidator) error
}

type UserRepository interface {
	FindByEmail(email string) (*entity.UserSchema, error)
	UpdateByEmail(email string, user entity.UserSchema) error
	CreateUsers(user entity.UserSchema) (*entity.UserSchema, error)
}
