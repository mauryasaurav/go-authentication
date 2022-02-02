package usecase

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mauryasaurav/go-authentication/domain/dto"
	"github.com/mauryasaurav/go-authentication/domain/entity"
	"github.com/mauryasaurav/go-authentication/domain/interfaces"
	"github.com/mauryasaurav/go-authentication/middleware/jwt"
	"golang.org/x/crypto/bcrypt"
)

type userUsecase struct {
	userRepo interfaces.UserRepository
}

func NewUserUsecase(userRepo interfaces.UserRepository) interfaces.UserUsecase {
	return &userUsecase{
		userRepo: userRepo,
	}
}

func (u *userUsecase) CreateUserHandler(ctx *gin.Context, req *dto.UserValidator) error {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	uuid, _ := uuid.NewUUID()

	user, err := u.userRepo.CreateUsers(entity.UserSchema{
		Id:        int64(uuid.ID()),
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  string(hashedPass),
	})

	if err != nil {
		return err
	}
	fmt.Println("user, tc======+>", user)
	return nil
}

func (u *userUsecase) UpdateUserHandler(ctx *gin.Context, req dto.UserUpdateValidator) error {
	err := u.userRepo.UpdateByEmail(req.Email, entity.UserSchema{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
	})

	if err != nil {
		return err
	}

	return nil
}

func (u *userUsecase) LoginUserHandler(ctx *gin.Context, req dto.UserLoginValidator) error {
	user, err := u.userRepo.FindByEmail(req.Email)
	if user == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User not found "})
		return err
	}
	valid := u.AuthenticationUser(ctx, req.Password, user.Password)

	if !valid {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Password don't match."})
		return nil
	}

	jwtToken := jwt.GenerateJWT(0)
	ctx.JSON(http.StatusOK, gin.H{"message": "user login successfully", "token": jwtToken})
	return nil
}

func (u *userUsecase) AuthenticationUser(ctx *gin.Context, oldPassword string, currentPassword string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(currentPassword), []byte(oldPassword)); err != nil {
		return false
	}
	return true
}
