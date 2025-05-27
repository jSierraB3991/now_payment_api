package nowpaymentsservice

import (
	"context"
	"encoding/json"

	jsierralibs "github.com/jSierraB3991/jsierra-libs"
	nowpaymentlibs "github.com/jSierraB3991/now_payment_api/domain/now_payment_libs"
	nowpaymentsmapper "github.com/jSierraB3991/now_payment_api/domain/now_payments_mapper"
	nowpaymentsrequest "github.com/jSierraB3991/now_payment_api/infrastructure/now_payments_request"
	nowpaymentsresponse "github.com/jSierraB3991/now_payment_api/infrastructure/now_payments_response"
)

func (s *NowPaymentService) CreateInvoice(ctx context.Context, req nowpaymentsrequest.CreateInvoiceRequest, userId uint, payCurrency string) (*nowpaymentsresponse.CreateInvoiceResponse, error) {
	jsonData, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	headers := []nowpaymentsrequest.HeaderRequest{{
		Key:   "x-api-key",
		Value: s.apiKey,
	}}
	var result nowpaymentsresponse.CreateInvoiceResponse
	err = s.httpClient.Post(s.apiUrl, nowpaymentlibs.CREATE_INVOICE_URL, jsonData, &result, headers)
	if err != nil {
		return nil, err
	}

	nowPaymentPaymentId, err := s.createPaymentByInvoice(ctx, result.ID, payCurrency)
	if err != nil {
		return nil, err
	}
	nowPaymentPaymentIdUint := jsierralibs.GetUNumberForString(*nowPaymentPaymentId)
	data := nowpaymentsmapper.GetCreateInvoiceByResponse(result, userId, nowPaymentPaymentIdUint)
	err = s.repository.SaveCreateInvoice(ctx, data)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
