package nowpaymentsrepository

import nowpaymentsmodel "github.com/jSierraB3991/now_payment_api/domain/now_payments_model"

func (repo *Repository) SaveCreatePayment(createPayment *nowpaymentsmodel.NowPaymentCreatePayment) error {
	return repo.db.Create(createPayment).Error
}

func (repo *Repository) SaveCreateInvoice(createInvoice *nowpaymentsmodel.NowPaymentCreateInvoice) error {
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
