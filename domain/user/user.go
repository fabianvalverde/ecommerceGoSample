package user

import (
	"github.com/fabianvalverde/ecommerceGoSample/model"
	"github.com/google/uuid"
)

// Puerto de entrada a user
type UseCase interface {
	Create(m *model.User) error
	GetByID(ID uuid.UUID) (model.User, error)
	GetByEmail(email string) (model.User, error)
	GetAll() (model.Users, error)
}

// Puerto de salida
type Storage interface {
	Create(m *model.User) error
	GetByID(ID uuid.UUID) (model.User, error)
	GetByEmail(email string) (model.User, error)
	GetAll() (model.Users, error)
}
