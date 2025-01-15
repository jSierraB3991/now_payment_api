package nowpaymentsrepositoryinterface

import nowpaymentsmodel "github.com/jSierraB3991/now_payment_api/domain/now_payments_model"

type PaymentRepositoryInterface interface {
	RunMigrations() error
	SaveCreatePayment(createPayment *nowpaymentsmodel.NowPaymentCreatePayment) error
	SaveCreateInvoice(createInvoice *nowpaymentsmodel.NowPaymentCreateInvoice) error
}
