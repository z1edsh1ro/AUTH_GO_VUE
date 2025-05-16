package router

import (
	"main/internal/adapter/http"
	"main/internal/adapter/repository"
	"main/internal/application/service"
	"main/internal/middleware"
	"time"

	"github.com/gin-contrib/cors"
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

	corsConfig := cors.Config{
		AllowOrigins:  []string{"http://localhost:5173"},
		AllowMethods:  []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:  []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders: []string{"Content-Length"},
		MaxAge:        12 * time.Hour,
	}

	r.Use(cors.New(corsConfig))

	api := r.Group("/api")

	auth := api.Group("/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
	}

	authUser := auth.Group("/user")
	authUser.Use(middleware.AuthMiddleware())
	{
		authUser.GET("/getAll", userHandler.GetAllUser)
	}

	return r
}
