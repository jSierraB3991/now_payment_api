package nowpaymentsservice

import (
	nowpaymentlibs "github.com/jSierraB3991/now_payment_api/domain/now_payment_libs"
	nowpaymentsresponse "github.com/jSierraB3991/now_payment_api/infrastructure/now_payments_response"
)

func (s *NowPaymentService) GetPaymentStatus(invoiceId string) (*nowpaymentsresponse.GetPaymentStatusResponse, error) {
	var result nowpaymentsresponse.GetPaymentStatusResponse
	err := s.httpClient.Get(s.apiUrl, nowpaymentlibs.GET_PAYMENT_DATA_URL+invoiceId, s.apiKey, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
