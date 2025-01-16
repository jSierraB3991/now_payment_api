package nowpaymentsservice

import (
	"log"

	nowpaymentsrepositoryinterface "github.com/jSierraB3991/now_payment_api/domain/now_payments_repository_interface"
	nowpaymentsclient "github.com/jSierraB3991/now_payment_api/infrastructure/now_payments_client"
	nowpaymentsrepository "github.com/jSierraB3991/now_payment_api/infrastructure/now_payments_repository"
	"gorm.io/gorm"
)

type NowPaymentService struct {
	repository nowpaymentsrepositoryinterface.PaymentRepositoryInterface
	apiKey     string
	apiUrl     string
	httpClient *nowpaymentsclient.HttpClient
}

func InitiateService(db *gorm.DB, nowPaymentApkiKey, apiUrl string) *NowPaymentService {
	repo := nowpaymentsrepository.InitiateRepo(db)

	return &NowPaymentService{
		repository: repo,
		apiKey:     nowPaymentApkiKey,
		apiUrl:     apiUrl,
		httpClient: nowpaymentsclient.NewHttpClient(),
	}
}

func (s *NowPaymentService) StartPayment() {
	err := s.repository.RunMigrations()
	if err != nil {
		log.Fatal(err)
	}
}
