package nowpaymentsservice

import (
	"log"
)

func (s *NowPaymentService) ValidateOnData(email, orderId string) (*string, error) {
	nowpyamentId, err := s.repository.GetInvoiceByUserAndOrder(email, orderId)
	if err != nil {
		return nil, err
	}
	return s.validateOnData(*nowpyamentId)
}

func (s *NowPaymentService) validateOnData(nowPaymentsInvoiceId string) (*string, error) {

	log.Printf("get invoice by invoice id %s", nowPaymentsInvoiceId)
	data, err := s.GetInvoiceStatus(nowPaymentsInvoiceId)
	if err != nil {
		return nil, err
	}

	return &data.PaymentStatus, err
}

func (s *NowPaymentService) Start(investPays []string) error {

	for _, v := range investPays {
		_, err := s.validateOnData(v)
		if err != nil {
			return err
		}
	}
	return nil
}
