package jwt

import (
	"time"

	jwt2 "github.com/dgrijalva/jwt-go"
	"github.com/mauryasaurav/go-authentication/utils/constants"
)

type Claims struct {
	UserID uint `json:"user_id"`
	jwt2.StandardClaims
}

func GenerateJWT(userID int) string {
	token := jwt2.New(jwt2.SigningMethodHS256)
	claims := token.Claims.(jwt2.MapClaims)
	claims["user_id"] = userID
	claims["expires_at"] = time.Now().Add(time.Minute * 5)

	secret := constants.SECRET_KEY
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return ""
	}
	return tokenString
}
