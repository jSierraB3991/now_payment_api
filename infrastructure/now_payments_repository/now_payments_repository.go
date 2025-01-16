package nowpaymentsrepository

import (
	nowpaymentlibs "github.com/jSierraB3991/now_payment_api/domain/now_payment_libs"
	nowpaymentsmodel "github.com/jSierraB3991/now_payment_api/domain/now_payments_model"
)

func (repo *Repository) SaveCreatePayment(createPayment *nowpaymentsmodel.NowPaymentCreatePayment) error {
	createPayment.Status = nowpaymentlibs.CREATE_INVOICE_STATUS_WAIT_PAY
	return repo.db.Create(createPayment).Error
}

func (repo *Repository) SaveCreateInvoice(createInvoice *nowpaymentsmodel.NowPaymentCreateInvoice) error {
	createInvoice.Status = nowpaymentlibs.CREATE_INVOICE_STATUS_WAIT_PAY
	return repo.db.Create(createInvoice).Error
}

func (repo *Repository) GetInvoicePagination(page *nowpaymentsmodel.Paggination, userId uint) ([]nowpaymentsmodel.NowPaymentCreateInvoice, error) {
	var result []nowpaymentsmodel.NowPaymentCreateInvoice
	params := []nowpaymentsmodel.PagginationParam{
		{
			Where: "user_id = ?",
			Data:  userId,
		},
	}
	preloads := []nowpaymentsmodel.PreloadParams{}

	err := repo.db.Scopes(repo.paginate_with_param(result, page, params, preloads)).Find(&result).Error
	return result, err
}

func (repo *Repository) UpdatePaymentIdInInvoiceId(invoiceId string, paymentId uint) error {
	var invoiceData nowpaymentsmodel.NowPaymentCreateInvoice
	err := repo.db.Where("now_payment_id = ?").Find(&invoiceData).Error
	if err != nil {
		return err
	}

	invoiceData.NowPaymentPaymentId = paymentId
	return repo.db.Save(&invoiceData).Error
}
