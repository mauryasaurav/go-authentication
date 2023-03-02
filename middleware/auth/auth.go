package auth

import (
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/mauryasaurav/go-authentication/utils/constants"
)

// AuthRequired is a simple middleware to check the session
func AuthRequired(c *gin.Context) {
	auth := c.Request.Header.Get("Authorization")
	if auth == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization is missing"})
		return
	}

	token := strings.Split(auth, "Bearer ")
	if len(token) == 0 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization is invalid"})
		return
	}

	if len(token) < 2 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token is invalid"})
		return
	}

	claims, valid := extractClaims(token[1])

	if !valid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	session := sessions.Default(c)

	role, ok := claims["role"].(string)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "role not found"})
		return
	}

	userId, ok := claims["user_id"].(string)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "user_id not found"})
		return
	}

	session.Set("role", role)
	session.Set("user_id", userId)
	session.Save()
	c.Next()
}

func extractClaims(tokenStr string) (jwt.MapClaims, bool) {
	hmacSecretString := constants.SECRET_KEY
	hmacSecret := []byte(hmacSecretString)
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// check token signing method etc
		return hmacSecret, nil
	})

	if err != nil {
		return nil, false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	} else {
		log.Printf("Invalid JWT Token")
		return nil, false
	}
}

func JSONMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Next()
	}
}
