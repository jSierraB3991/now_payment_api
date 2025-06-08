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
	apiUrl     string
	httpClient *nowpaymentsclient.HttpClient
}

func InitiateService(db *gorm.DB, apiUrl string) *NowPaymentService {
	repo := nowpaymentsrepository.InitiateRepo(db)

	return &NowPaymentService{
		repository: repo,
		apiUrl:     apiUrl,
		httpClient: nowpaymentsclient.NewHttpClient(),
	}
}

func (s *NowPaymentService) StartPaymentWithOutSchema() {
	err := s.repository.RunMigrations([]string{"public"})
	if err != nil {
		log.Fatal(err)
	}
}

func (s *NowPaymentService) StartPayment(schemas []string) {
	err := s.repository.RunMigrations(schemas)
	if err != nil {
		log.Fatal(err)
	}
}
