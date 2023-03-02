package jwt

import (
	"time"

	jwt2 "github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/mauryasaurav/go-authentication/utils/constants"
)

type Claims struct {
	UserID uint `json:"user_id"`
	jwt2.StandardClaims
}

func GenerateJWTToken(userID uuid.UUID, role string) string {
	token := jwt2.New(jwt2.SigningMethodHS256)
	claims := token.Claims.(jwt2.MapClaims)
	claims["user_id"] = userID
	claims["role"] = role
	claims["expires_at"] = time.Now().Add(time.Minute * 100)

	tokenString, err := token.SignedString([]byte(constants.SECRET_KEY))
	if err != nil {
		return ""
	}

	return tokenString
}
