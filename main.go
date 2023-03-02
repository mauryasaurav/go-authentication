package main

import (
	"fmt"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	handler "github.com/mauryasaurav/go-authentication/api/http"
	"github.com/mauryasaurav/go-authentication/api/repozitory"
	"github.com/mauryasaurav/go-authentication/api/usecase"
	"github.com/mauryasaurav/go-authentication/domain/entity"
	"github.com/mauryasaurav/go-authentication/middleware/auth"
	"github.com/mauryasaurav/go-authentication/utils/constants"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	dbConn, err := gorm.Open(postgres.Open(constants.DB_URL), &gorm.Config{})

	if err != nil {
		fmt.Println(err.Error())
		panic("dbConn not connected")
	}

	// dbConn.Debug()
	dbConn.AutoMigrate(&entity.CategorySchema{})
	dbConn.AutoMigrate(&entity.UserSchema{})
	dbConn.AutoMigrate(&entity.QuestionSchema{})

	route := gin.Default()

	route.Use(auth.JSONMiddleware())

	store := cookie.NewStore([]byte("secret"))
	route.Use(sessions.Sessions("mysession", store))

	api := route.Group("/api/")

	// public
	publicRoute := api.Group("/public")
	publicRepo := repozitory.NewPublicRepository(dbConn)
	publicUsecase := usecase.NewPublicUsecase(publicRepo)
	handler.NewPublicHandler(publicRoute, publicUsecase)

	// user
	userRoute := api.Group("/user")
	userRepo := repozitory.NewUserRepository(dbConn)
	userUsecase := usecase.NewUserUsecase(userRepo)
	handler.NewUserHandler(userRoute, userUsecase)

	// private
	private := api.Group("/private")
	private.Use(auth.AuthRequired)
	{
		// Category
		categoryRepo := repozitory.NewCategoryRepository(dbConn)
		categoryUsecase := usecase.NewCategoryUsecase(categoryRepo)
		handler.NewCategoryHandler(private, categoryUsecase)

		// Question
		questionRepo := repozitory.NewQuestionRepository(dbConn)
		questionUsecase := usecase.NewQuestionUsecase(questionRepo)
		handler.NewQuestionHandler(private, questionUsecase)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	route.Run(":" + port)
}
