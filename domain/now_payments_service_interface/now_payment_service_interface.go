package nowpaymentsserviceinterface

import (
	nowpaymentsrequest "github.com/jSierraB3991/now_payment_api/infrastructure/now_payments_request"
	nowpaymentsresponse "github.com/jSierraB3991/now_payment_api/infrastructure/now_payments_response"
)

type NowPaymentServiceInterface interface {
	CreatePayment(req nowpaymentsrequest.CreatePaymentRequest, userId uint) (*nowpaymentsresponse.CreatePaymentResponse, error)
	CreateInvoice(req nowpaymentsrequest.CreateInvoiceRequest, userId uint) (*nowpaymentsresponse.CreateInvoiceResponse, error)
}
