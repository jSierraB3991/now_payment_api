package nowpaymentsrepository

import (
	nowpaymentserrors "github.com/jSierraB3991/now_payment_api/domain/now_payments_errors"
	nowpaymentsmodel "github.com/jSierraB3991/now_payment_api/domain/now_payments_model"
)

func (repo Repository) GetInvoiceByUserAndOrder(userEmail, orderId string) (*string, error) {
	var result []string
	modelDb := repo.db.Select("now_payment_id").Model(&nowpaymentsmodel.NowPaymentCreateInvoice{})
	err := modelDb.Where("order_id  = ?", orderId).Find(&result).Error
	if err != nil {
		return nil, err
	}
	if len(result) != 1 {
		return nil, nowpaymentserrors.InvalidInvestOrderIdError{OrderId: orderId}
	}
	return &result[0], nil
}
