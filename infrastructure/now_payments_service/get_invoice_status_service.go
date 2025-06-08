package nowpaymentsservice

import (
	"context"

	nowpaymentlibs "github.com/jSierraB3991/now_payment_api/domain/now_payment_libs"
	nowpaymentserrors "github.com/jSierraB3991/now_payment_api/domain/now_payments_errors"
	nowpaymentsrequest "github.com/jSierraB3991/now_payment_api/infrastructure/now_payments_request"
	nowpaymentsresponse "github.com/jSierraB3991/now_payment_api/infrastructure/now_payments_response"
)

func (s *NowPaymentService) GetInvoiceStatus(ctx context.Context, invoiceId, userName, password, apiKey string) (*nowpaymentsresponse.GetPaymentStatusResponse, error) {
	token, err := s.AuthPaymentService(userName, password, apiKey)
	if err != nil {
		return nil, err
	}

	headers := []nowpaymentsrequest.HeaderRequest{{
		Key:   "Authorization",
		Value: "Bearer " + *token,
	}, {
		Key:   "x-api-key",
		Value: apiKey,
	}}

	var resultPaymentsData nowpaymentsresponse.GetPaymensDataStatusResponse
	err = s.httpClient.Get(s.apiUrl, nowpaymentlibs.GET_PAYMENT_DATA_URL+"?invoiceId="+invoiceId, &resultPaymentsData, headers)
	if err != nil {
		return nil, err
	}

	if len(resultPaymentsData.Data) <= 0 {
		return nil, nowpaymentserrors.InvalidInvoicePaymentError{}
	}
	result := s.GetNotWaitingResponse(resultPaymentsData.Data)

	if result.PaymentId != 0 {
		status := func(status string) nowpaymentlibs.CreateInvoiceStatus {
			switch status {
			case "waiting":
				return nowpaymentlibs.CREATE_INVOICE_STATUS_WAIT_PAY
			case "confirming":
				return nowpaymentlibs.CREATE_INVOICE_STATUS_CONFIRMING
			case "confirmed":
			case "finished":
				return nowpaymentlibs.CREATE_INVOICE_STATUS_PAY
			case "failed":
			case "expired":
				return nowpaymentlibs.CREATE_INVOICE_STATUS_CANCELL
			}
			return nowpaymentlibs.CREATE_INVOICE_STATUS_WAIT_PAY
		}(result.PaymentStatus)
		err = s.repository.UpdatePaymentIdInInvoiceId(ctx, invoiceId, result.PaymentId, status)
		if err != nil {
			return nil, err
		}
	}

	return &result, nil
}

func (NowPaymentService) GetNotWaitingResponse(data []nowpaymentsresponse.GetPaymentStatusResponse) nowpaymentsresponse.GetPaymentStatusResponse {
	if len(data) == 1 {
		return data[0]
	}

	var preResult []nowpaymentsresponse.GetPaymentStatusResponse
	for _, v := range data {
		if v.PaymentStatus != "waiting" {
			preResult = append(preResult, v)
		}
	}

	if len(preResult) == 0 {
		return data[0]
	} else if len(preResult) == 1 {
		return preResult[0]
	}

	var result nowpaymentsresponse.GetPaymentStatusResponse
	for _, v := range preResult {
		if v.PaymentStatus == "finished" || v.PaymentStatus == "confirmed" {
			result = v
			break
		}
	}
	if result.PaymentId == 0 {
		return preResult[1]
	}
	return result
}
