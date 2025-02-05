package nowpaymentsrepositoryinterface

import (
	nowpaymentlibs "github.com/jSierraB3991/now_payment_api/domain/now_payment_libs"
	nowpaymentsmodel "github.com/jSierraB3991/now_payment_api/domain/now_payments_model"
)

type PaymentRepositoryInterface interface {
	RunMigrations() error
	SaveCreatePayment(createPayment *nowpaymentsmodel.NowPaymentCreatePayment) error
	SaveCreateInvoice(createInvoice *nowpaymentsmodel.NowPaymentCreateInvoice) error
	GetInvoicePagination(page *nowpaymentsmodel.Paggination, userId uint) ([]nowpaymentsmodel.NowPaymentCreateInvoice, error)
	UpdatePaymentIdInInvoiceId(invoiceId string, paymentId uint, status nowpaymentlibs.CreateInvoiceStatus) error
	GetInvoiceByUserAndOrder(userEmail, orderId string) (*string, error)
}
