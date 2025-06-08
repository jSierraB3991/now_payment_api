package nowpaymentsservice

import (
	"context"
	"encoding/json"

	nowpaymentlibs "github.com/jSierraB3991/now_payment_api/domain/now_payment_libs"
	nowpaymentsmapper "github.com/jSierraB3991/now_payment_api/domain/now_payments_mapper"
	nowpaymentsrequest "github.com/jSierraB3991/now_payment_api/infrastructure/now_payments_request"
	nowpaymentsresponse "github.com/jSierraB3991/now_payment_api/infrastructure/now_payments_response"
)

func (s *NowPaymentService) CreatePayment(ctx context.Context, req nowpaymentsrequest.CreatePaymentRequest, userId uint, apiKey string) (*nowpaymentsresponse.CreatePaymentResponse, error) {
	jsonData, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	headers := []nowpaymentsrequest.HeaderRequest{{
		Key:   "x-api-key",
		Value: apiKey,
	}}
	var result nowpaymentsresponse.CreatePaymentResponse
	err = s.httpClient.Post(s.apiUrl, nowpaymentlibs.CREATE_PAYMENT_URL, jsonData, &result, headers)
	if err != nil {
		return nil, err
	}
	data := nowpaymentsmapper.GetCreatePaymentByResponse(result, userId)
	err = s.repository.SaveCreatePayment(ctx, &data)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
