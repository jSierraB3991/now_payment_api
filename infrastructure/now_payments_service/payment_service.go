package nowpaymentsservice

import (
	nowpaymentsrepositoryinterface "github.com/jSierraB3991/now_payment_api/domain/now_payments_repository_interface"
	nowpaymentsrepository "github.com/jSierraB3991/now_payment_api/infrastructure/now_payments_repository"
	"gorm.io/gorm"
)

type NowPaymentService struct {
	repository nowpaymentsrepositoryinterface.PaymentRepositoryInterface
	apiKey     string
	apiUrl     string
}

func InitiateService(db *gorm.DB, nowPaymentApkiKey string) *NowPaymentService {
	repo := nowpaymentsrepository.InitiateRepo(db)

	return &NowPaymentService{
		repository: repo,
		apiKey:     nowPaymentApkiKey,
		apiUrl:     "https://api.nowpayments.io/v1",
	}
}

func (s *NowPaymentService) StartPayment() {
	// Start payment logic
}
