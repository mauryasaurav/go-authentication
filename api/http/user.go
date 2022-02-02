package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mauryasaurav/go-authentication/domain/dto"
	"github.com/mauryasaurav/go-authentication/domain/interfaces"
)

type userHandler struct {
	userUsecase interfaces.UserUsecase
}

func NewUserHandler(route *gin.RouterGroup, u interfaces.UserUsecase) {
	handler := userHandler{userUsecase: u}
	route.POST("/create", handler.CreateUser)
	route.PUT("/update", handler.UpdateUser)
	route.POST("/login", handler.LoginUser)
}

func (h *userHandler) CreateUser(ctx *gin.Context) {
	created := new(dto.UserValidator)
	if err := ctx.Bind(created); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.userUsecase.CreateUserHandler(ctx, created)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

func (h *userHandler) UpdateUser(ctx *gin.Context) {
	updated := new(dto.UserUpdateValidator)
	if err := ctx.Bind(&updated); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.userUsecase.UpdateUserHandler(ctx, *updated)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusBadRequest, gin.H{"message": "Update user successfully"})
}

func (h *userHandler) LoginUser(ctx *gin.Context) {
	req := new(dto.UserLoginValidator)
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.userUsecase.LoginUserHandler(ctx, *req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}
