package helper

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	jwt2 "github.com/mauryasaurav/go-authentication/middleware/jwt"
	"github.com/patrickmn/go-cache"
)

var Cache = cache.New(30*time.Second, 1*time.Minute)

func GetUserIDFromCtx(ctx *gin.Context) (uint, error) {
	token, ok := ctx.Get("user").(*jwt.Token)
	if !ok || token == nil {
		return 0, errors.New("error getting user_id from ctx")
	}

	claims, ok := token.Claims.(*jwt2.Claims)
	if !ok || claims == nil {
		return 0, errors.New("error getting user_id from ctx")
	}

	return claims.UserID, nil
}
