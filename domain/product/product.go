package product

import (
	"github.com/fabianvalverde/ecommerceGoSample/model"

	"github.com/google/uuid"
)

// Puerto de entrada a product
type UseCase interface {
	Create(m *model.Product) error
	Update(m *model.Product) error
	Delete(ID uuid.UUID) error

	GetByID(ID uuid.UUID) (model.Product, error)
	GetAll() (model.Products, error)
}

// Puerto de salida
type Storage interface {
	Create(m *model.Product) error
	Update(m *model.Product) error
	Delete(ID uuid.UUID) error

	GetByID(ID uuid.UUID) (model.Product, error)
	GetAll() (model.Products, error)
}
