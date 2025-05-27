package nowpaymentsservice

import (
	"context"

	nowpaymentlibs "github.com/jSierraB3991/now_payment_api/domain/now_payment_libs"
	nowpaymentsrequest "github.com/jSierraB3991/now_payment_api/infrastructure/now_payments_request"
	nowpaymentsresponse "github.com/jSierraB3991/now_payment_api/infrastructure/now_payments_response"
)

func (s *NowPaymentService) GetPaymentStatus(ctx context.Context, invoiceId string) (*nowpaymentsresponse.GetPaymentStatusResponse, error) {

	headers := []nowpaymentsrequest.HeaderRequest{{
		Key:   "x-api-key",
		Value: s.apiKey,
	}}

	var result nowpaymentsresponse.GetPaymentStatusResponse
	err := s.httpClient.Get(s.apiUrl, nowpaymentlibs.GET_PAYMENT_DATA_URL+invoiceId, &result, headers)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
