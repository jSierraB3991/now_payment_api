package nowpaymentsservice

import (
	nowpaymentlibs "github.com/jSierraB3991/now_payment_api/domain/now_payment_libs"
	nowpaymentserrors "github.com/jSierraB3991/now_payment_api/domain/now_payments_errors"
	nowpaymentsrequest "github.com/jSierraB3991/now_payment_api/infrastructure/now_payments_request"
	nowpaymentsresponse "github.com/jSierraB3991/now_payment_api/infrastructure/now_payments_response"
)

func (s *NowPaymentService) GetInvoiceStatus(invoiceId string) (*nowpaymentsresponse.GetPaymentStatusResponse, error) {
	token, err := s.AuthPaymentService()
	if err != nil {
		return nil, err
	}

	headers := []nowpaymentsrequest.HeaderRequest{{
		Key:   "Authorization",
		Value: "Bearer " + *token,
	}, {
		Key:   "x-api-key",
		Value: s.apiKey,
	}}

	var resultPaymentsData nowpaymentsresponse.GetPaymensDataStatusResponse
	err = s.httpClient.Get(s.apiUrl, nowpaymentlibs.GET_PAYMENT_DATA_URL+"?invoiceId="+invoiceId, &resultPaymentsData, headers)
	if err != nil {
		return nil, err
	}

	if len(resultPaymentsData.Data) != 1 {
		return nil, nowpaymentserrors.InvalidInvoicePaymentError{}
	}
	result := resultPaymentsData.Data[0]

	if result.PaymentId != 0 {
		err = s.repository.UpdatePaymentIdInInvoiceId(invoiceId, result.PaymentId)
		if err != nil {
			return nil, err
		}
	}

	return &result, nil
}
