package http

import (
	"main/internal/application/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	Service service.UserService
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{Service: service}
}

func (handler *UserHandler) GetAllUser(context *gin.Context) {
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
