package nowpaymentsrepository

import (
	"context"

	jsierralibs "github.com/jSierraB3991/jsierra-libs"
	nowpaymentlibs "github.com/jSierraB3991/now_payment_api/domain/now_payment_libs"
	nowpaymentsmodel "github.com/jSierraB3991/now_payment_api/domain/now_payments_model"
)

func (repo *Repository) SaveCreatePayment(ctx context.Context, createPayment *nowpaymentsmodel.NowPaymentCreatePayment) error {
	db, err := repo.GetDb(ctx)
	if err != nil {
		return err
	}
	createPayment.Status = nowpaymentlibs.CREATE_INVOICE_STATUS_WAIT_PAY
	return db.Create(createPayment).Error
}

func (repo *Repository) SaveCreateInvoice(ctx context.Context, createInvoice *nowpaymentsmodel.NowPaymentCreateInvoice) error {
	db, err := repo.GetDb(ctx)
	if err != nil {
		return err
	}
	createInvoice.Status = nowpaymentlibs.CREATE_INVOICE_STATUS_WAIT_PAY
	return db.Create(createInvoice).Error
}

func (repo *Repository) GetInvoiceByOrderId(ctx context.Context, orderId string) (*nowpaymentsmodel.NowPaymentCreateInvoice, error) {
	db, err := repo.GetDb(ctx)
	if err != nil {
		return nil, err
	}

	var invoice nowpaymentsmodel.NowPaymentCreateInvoice
	err = db.Where("order_id = ?", orderId).First(&invoice).Error
	if err != nil {
		return nil, err
	}
	return &invoice, nil
}

func (repo *Repository) GetInvoicePagination(ctx context.Context, page *jsierralibs.Paggination, userId uint) ([]nowpaymentsmodel.NowPaymentCreateInvoice, error) {
	db, err := repo.GetDb(ctx)
	if err != nil {
		return nil, err
	}

	var result []nowpaymentsmodel.NowPaymentCreateInvoice
	params := []jsierralibs.PagginationParam{
		{
			Where: "user_id = ?",
			Data:  []interface{}{userId},
		},
	}
	preloads := []jsierralibs.PreloadParams{}

	err = db.Scopes(repo.paginate_with_param(ctx, result, page, params, preloads)).Find(&result).Error
	return result, err
}

func (repo *Repository) UpdatePaymentIdInInvoiceId(ctx context.Context, invoiceId string, paymentId uint, status nowpaymentlibs.CreateInvoiceStatus) error {
	db, err := repo.GetDb(ctx)
	if err != nil {
		return err
	}

	var invoiceData nowpaymentsmodel.NowPaymentCreateInvoice
	err = db.Where("now_payment_id = ?", invoiceId).Find(&invoiceData).Error
	if err != nil {
		return err
	}
	if invoiceData.NowPaymentPaymentId == paymentId && invoiceData.Status == status {
		return nil
	}

	invoiceData.NowPaymentPaymentId = paymentId
	invoiceData.Status = status
	return repo.db.Save(&invoiceData).Error
}
