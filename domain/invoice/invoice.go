package invoice

import (
	"github.com/fabianvalverde/ecommerceGoSample/model"
	"github.com/google/uuid"
)

// Puerto de entrada invoice
type UseCase interface {
	Create(m *model.PurchaseOrder) error
	GetByUserID(userID uuid.UUID) (model.InvoicesReport, error)
	GetAll() (model.InvoicesReport, error)
}

// Puerto de salida
type Storage interface {
	Create(m *model.Invoice, ms model.InvoiceDetails) error
}

// Puerto de salida
type StorageInvoiceDetailReport interface {
	HeadByInvoiceID(ID uuid.UUID) (model.InvoiceReport, error)
	HeadsByUserID(userID uuid.UUID) (model.InvoicesReport, error)
	AllHead() (model.InvoicesReport, error)
	AllDetailsByInvoiceID(ID uuid.UUID) (model.InvoiceDetailsReports, error)
}
