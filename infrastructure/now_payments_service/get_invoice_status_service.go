package nowpaymentsservice

import (
	nowpaymentsresponse "github.com/jSierraB3991/now_payment_api/infrastructure/now_payments_response"
)

func (s *NowPaymentService) GetInvoiceStatus(invoiceId string) (*nowpaymentsresponse.GetPaymentStatusResponse, error) {
	var result nowpaymentsresponse.GetPaymentStatusResponse
	err := s.httpClient.Get(s.apiUrl, ""+invoiceId, s.apiKey, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
