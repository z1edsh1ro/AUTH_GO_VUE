package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		log.Printf("test")
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
			context.Abort()
			return
		}

		_, ok := token.Claims.(jwt.MapClaims)

		if !ok {
			context.JSON(http.StatusUnauthorized, gin.H{
				"status":  "ERROR",
				"message": "Invalid token claims",
			})
			context.Abort()
			return
		}

		context.Next()
	}
}
