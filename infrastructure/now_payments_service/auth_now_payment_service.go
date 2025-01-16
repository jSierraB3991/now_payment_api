package nowpaymentsservice

import (
	"encoding/json"

	nowpaymentlibs "github.com/jSierraB3991/now_payment_api/domain/now_payment_libs"
	nowpaymentsrequest "github.com/jSierraB3991/now_payment_api/infrastructure/now_payments_request"
	nowpaymentsresponse "github.com/jSierraB3991/now_payment_api/infrastructure/now_payments_response"
)

func (s *NowPaymentService) AuthPaymentService() (*string, error) {
	var result nowpaymentsresponse.AuthResponse
	req := nowpaymentsrequest.AuthRequest{
		Email:    s.UserName,
		Password: s.Password,
	}
	jsonData, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	headers := []nowpaymentsrequest.HeaderRequest{{
		Key:   "x-api-key",
		Value: s.apiKey,
	}}
	err = s.httpClient.Post(s.apiUrl, nowpaymentlibs.AUTH_URL, jsonData, &result, headers)
	if err != nil {
		return nil, err
	}
	return &result.Token, nil
}
