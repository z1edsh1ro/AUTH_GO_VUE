package router

import (
	"main/internal/adapter/http"
	"main/internal/adapter/repository"
	"main/internal/application/service"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func New(mongoClient *mongo.Client) *gin.Engine {
	r := gin.Default()

	userRepository := repository.NewUserRepository(mongoClient)
	authService := service.AuthService{Port: userRepository}
	userService := service.UserService{Port: userRepository}

	authHandler := http.NewAuthHandler(authService)
	userHandler := http.NewUserHandler(userService)

	apiAuth := r.Group("/api/auth")
	{
		apiAuth.POST("/register", authHandler.Register)
		apiAuth.POST("/login", authHandler.Login)
	}

	apiAuthUser := r.Group("/api/auth/user")
	{
		apiAuthUser.GET("/", userHandler.GetAllUser)
	}

	return r
}
