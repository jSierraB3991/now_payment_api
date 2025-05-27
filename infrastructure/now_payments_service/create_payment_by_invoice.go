package nowpaymentsservice

import (
	"context"
	"encoding/json"

	nowpaymentlibs "github.com/jSierraB3991/now_payment_api/domain/now_payment_libs"
	nowpaymentsrequest "github.com/jSierraB3991/now_payment_api/infrastructure/now_payments_request"
	nowpaymentsresponse "github.com/jSierraB3991/now_payment_api/infrastructure/now_payments_response"
)

func (s *NowPaymentService) createPaymentByInvoice(ctx context.Context, invoiceId, payCurrency string) (*string, error) {
	req := nowpaymentsrequest.CreatePaymentByInvoiceRequest{
		Iid:         invoiceId,
		PayCurrency: payCurrency,
	}

	jsonData, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	headers := []nowpaymentsrequest.HeaderRequest{{
		Key:   "x-api-key",
		Value: s.apiKey,
	}}
	var result nowpaymentsresponse.CreatePaymentByInvoiceResponse
	err = s.httpClient.Post(s.apiUrl, nowpaymentlibs.CREATE_PAYMENT_BY_INVOICE, jsonData, &result, headers)
	if err != nil {
		return nil, err
	}
	return &result.PaymentId, nil
}
