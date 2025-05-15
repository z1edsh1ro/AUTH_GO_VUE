package router

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func New(mongoClient *mongo.Client) *gin.Engine {
	r := gin.Default()

	// todoRepository := repository.NewAuthRepository()
	// todoService := service.TodoService{Port: todoRepository}
	// todoHandler := http.NewTodoHandler(todoService)

	apiAuth := r.Group("/api/auth")
	{
		apiAuth.POST("/register", authHandler.Create)
		apiAuth.POST("/login", authHandler.Update)
	}

	apiAuthUser := r.Group("/api/auth/user")
	{
		apiAuthUser.GET("/register", userHandler.Create)
	}

	return r
}
