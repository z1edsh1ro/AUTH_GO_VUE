package port

import (
	"main/internal/domain/model"
)

type UserRepositoryInterface interface {
	GetAll() ([]model.User, error)
	GetByEmail(email string) (model.User, error)
	Create(entry model.User) error
}
