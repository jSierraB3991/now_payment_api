package nowpaymentsrepository

import nowpaymentsmodel "github.com/jSierraB3991/now_payment_api/domain/now_payments_model"

func (repo *Repository) RunMigrations() error {
	return repo.runAutoMigrate()
}

func (repo *Repository) runAutoMigrate() error {
	return repo.db.AutoMigrate(
		&nowpaymentsmodel.NowPaymentCreatePayment{},
	)
}
