package paypal

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"

	"github.com/fabianvalverde/ecommerceGoSample/domain/invoice"
	"github.com/fabianvalverde/ecommerceGoSample/domain/paypal"
	"github.com/fabianvalverde/ecommerceGoSample/domain/purchaseorder"

	storageInvoice "github.com/fabianvalverde/ecommerceGoSample/infrastructure/postgres/invoice"
	storagePurchaseOrder "github.com/fabianvalverde/ecommerceGoSample/infrastructure/postgres/purchaseorder"
)

func NewRouter(e *echo.Echo, dbPool *pgxpool.Pool) {
	h := buildHandler(dbPool)

	publicRoutes(e, h)
}

func buildHandler(dbPool *pgxpool.Pool) handler {
	purchaseOrderUseCase := purchaseorder.New(storagePurchaseOrder.New(dbPool))
	invoiceUseCase := invoice.New(storageInvoice.New(dbPool), nil)
	useCase := paypal.New(purchaseOrderUseCase, invoiceUseCase)

	return newHandler(useCase)
}

// publicRoutes handle the routes that not requires a validation of any kind to be use
func publicRoutes(e *echo.Echo, h handler) {
	route := e.Group("/api/v1/public/paypal")

	route.POST("", h.Webhook)
}
