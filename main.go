package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	handler "github.com/mauryasaurav/go-authentication/api/http"
	"github.com/mauryasaurav/go-authentication/api/repozitory"
	"github.com/mauryasaurav/go-authentication/api/usecase"
	"github.com/mauryasaurav/go-authentication/utils/constants"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Auth struct {
	Username string `json:"user_name"`
	Password string `json:"password"`
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

// AuthRequired is a simple middleware to check the session
func AuthRequired(c *gin.Context) {
	auth := c.Request.Header.Get("Authorization")
	token := strings.Split(auth, "Bearer ")
	data, valid := extractClaims(token[1])
	fmt.Println("data=====+>", valid, data)
	if !valid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	c.Next()
}

func JSONMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Next()
	}
}

func main() {

	dbConn, err := gorm.Open(postgres.Open(constants.DB_URL), &gorm.Config{})

	if err != nil {
		fmt.Println(err.Error())
		panic("dbConn not connected")
	}

	dbConn.AutoMigrate(&Auth{})

	route := gin.Default()

	route.Use(JSONMiddleware())

	api := route.Group("/api/")

	// repozitory setup
	userRepo := repozitory.NewUserRepository(dbConn)

	// User route
	userRoute := api.Group("/user")
	private := api.Group("/private")

	private.Use(AuthRequired)
	{
		private.POST("/me", handler.TestHere)
	}

	userUsecase := usecase.NewUserUsecase(userRepo)
	handler.NewUserHandler(userRoute, userUsecase)

	route.Run(":8001")
}
