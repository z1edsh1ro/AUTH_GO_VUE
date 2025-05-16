package service

import (
	"log"
	"main/internal/domain/model"
	"main/internal/domain/port"
)

type UserService struct {
	Port port.UserRepositoryInterface
}

func (service *UserService) GetAllUser() ([]model.User, error) {
	users, err := service.Port.GetAll()

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return users, nil
}
