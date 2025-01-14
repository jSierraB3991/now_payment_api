package nowpaymentsrepository

import nowpaymentsmodel "github.com/jSierraB3991/now_payment_api/domain/now_payments_model"

func (repo *Repository) SaveCreatePayment(createPayment *nowpaymentsmodel.NowPaymentCreatePayment) error {
	return repo.db.Create(createPayment).Error
}
