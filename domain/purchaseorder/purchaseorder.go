package purchaseorder

import (
	"github.com/fabianvalverde/ecommerceGoSample/model"

	"github.com/google/uuid"
)

// Puerto de entrada a PurchaseOrder
type UseCase interface {
	Create(m *model.PurchaseOrder) error

	GetByID(ID uuid.UUID) (model.PurchaseOrder, error)
}

// Puerto de salida
type Storage interface {
	Create(m *model.PurchaseOrder) error

	GetByID(ID uuid.UUID) (model.PurchaseOrder, error)
}
