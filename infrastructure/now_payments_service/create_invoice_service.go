package nowpaymentsservice

import (
	"encoding/json"

	jsierralibs "github.com/jSierraB3991/jsierra-libs"
	nowpaymentlibs "github.com/jSierraB3991/now_payment_api/domain/now_payment_libs"
	nowpaymentsmapper "github.com/jSierraB3991/now_payment_api/domain/now_payments_mapper"
	nowpaymentsrequest "github.com/jSierraB3991/now_payment_api/infrastructure/now_payments_request"
	nowpaymentsresponse "github.com/jSierraB3991/now_payment_api/infrastructure/now_payments_response"
)

func (s *NowPaymentService) CreateInvoice(req nowpaymentsrequest.CreateInvoiceRequest, userId uint) (*nowpaymentsresponse.CreateInvoiceResponse, error) {
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

	nowPaymentPaymentId, err := s.createPaymentByInvoice(result.ID)
	if err != nil {
		return nil, err
	}
	nowPaymentPaymentIdUint := jsierralibs.GetUNumberForString(*nowPaymentPaymentId)
	data := nowpaymentsmapper.GetCreateInvoiceByResponse(result, userId, nowPaymentPaymentIdUint)
	err = s.repository.SaveCreateInvoice(data)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
