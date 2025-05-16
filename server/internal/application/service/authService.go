package service

import (
	"fmt"
	"log"
	"main/internal/domain/dto"
	"main/internal/domain/model"
	"main/internal/domain/port"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	Port port.UserRepositoryInterface
}

func (service *AuthService) Register(entry dto.RegisterRequestDTO) error {
	user, err := service.Port.GetByEmail(entry.Email)

	if (user != model.User{}) {
		log.Println("ERROR ACCOUNT EXIST")
		return fmt.Errorf("ERROR ACCOUNT EXIST")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(entry.Password), 10)

	if err != nil {
		log.Println(err)
		return err
	}

	createUser := model.User{
		Name:     entry.Name,
		Email:    entry.Email,
		Password: string(hash),
	}

	err = service.Port.Create(createUser)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (service *AuthService) Login(entry dto.LoginRequestDTO) (dto.UserResponseDTO, string, error) {
	user, err := service.Port.GetByEmail(entry.Email)

	if (user == model.User{}) {
		log.Println("ERROR USER NOT FOUND")
		return dto.UserResponseDTO{}, "", fmt.Errorf("ERROR USER NOT FOUND")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(entry.Password))

	if err != nil {
		log.Println("ERROR PASSWORD WRONG")
		return dto.UserResponseDTO{}, "", fmt.Errorf("ERROR PASSWORD WRONG")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Email,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		log.Println(err)
		return dto.UserResponseDTO{}, "", err
	}

	userData := dto.UserResponseDTO{
		Name:  user.Name,
		Email: user.Email,
	}

	return userData, tokenString, nil
}
