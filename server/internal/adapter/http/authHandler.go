package http

import (
	"main/internal/application/service"
	"main/internal/domain/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	Service service.AuthService
}

func NewAuthHandler(service service.AuthService) *AuthHandler {
	return &AuthHandler{Service: service}
}

func (handler *AuthHandler) Register(context *gin.Context) {
	var user model.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  "ERROR",
			"message": err.Error(),
		})
		return
	}

	err = handler.Service.Register(user)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status":  "ERROR",
			"message": err.Error(),
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"status":  "CREATED",
		"message": "REGISTER USER SUCCESS",
	})
}

func (handler *AuthHandler) Login(context *gin.Context) {
	var user model.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  "ERROR",
			"message": err.Error(),
		})
		return
	}

	userDto, jwt, err := handler.Service.Login(user)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status":  "ERROR",
			"message": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status": "SUCCESS",
		"jwt":    jwt,
		"data":   userDto,
	})
}
