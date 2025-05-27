package nowpaymentsrepositoryinterface

import (
	"context"

	jsierralibs "github.com/jSierraB3991/jsierra-libs"
	nowpaymentlibs "github.com/jSierraB3991/now_payment_api/domain/now_payment_libs"
	nowpaymentsmodel "github.com/jSierraB3991/now_payment_api/domain/now_payments_model"
)

type PaymentRepositoryInterface interface {
	RunMigrations(schemas []string) error
	SaveCreatePayment(ctx context.Context, createPayment *nowpaymentsmodel.NowPaymentCreatePayment) error
	SaveCreateInvoice(ctx context.Context, createInvoice *nowpaymentsmodel.NowPaymentCreateInvoice) error
	GetInvoicePagination(ctx context.Context, page *jsierralibs.Paggination, userId uint) ([]nowpaymentsmodel.NowPaymentCreateInvoice, error)
	UpdatePaymentIdInInvoiceId(ctx context.Context, invoiceId string, paymentId uint, status nowpaymentlibs.CreateInvoiceStatus) error
}
