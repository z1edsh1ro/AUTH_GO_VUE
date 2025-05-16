package http

import (
	"fmt"
	"main/internal/application/service"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type UserHandler struct {
	Service service.UserService
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{Service: service}
}

func (handler *UserHandler) GetAllUser(context *gin.Context) {
	header := context.Request.Header.Get("Authorization")

	tokenString := strings.SplitAfter(header, " ")

	secret := []byte(os.Getenv("SECRET"))

	token, err := jwt.Parse(tokenString[1], func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
		}

		return secret, nil
	})

	if err != nil || !token.Valid {
		context.JSON(http.StatusUnauthorized, gin.H{
			"status":  "ERROR",
			"message": "Invalid or expired token",
		})
		return
	}

	_, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		context.JSON(http.StatusUnauthorized, gin.H{
			"status":  "ERROR",
			"message": "Invalid token claims",
		})
		return
	}

	users, err := handler.Service.GetAllUser()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  "ERROR",
			"message": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status": "SUCCESS",
		"data":   users,
	})
}
