package nowpaymentsserviceinterface

import (
	nowpaymentsmodel "github.com/jSierraB3991/now_payment_api/domain/now_payments_model"
	nowpaymentsrequest "github.com/jSierraB3991/now_payment_api/infrastructure/now_payments_request"
	nowpaymentsresponse "github.com/jSierraB3991/now_payment_api/infrastructure/now_payments_response"
)

type NowPaymentServiceInterface interface {
	CreatePayment(req nowpaymentsrequest.CreatePaymentRequest, userId uint) (*nowpaymentsresponse.CreatePaymentResponse, error)
	CreateInvoice(req nowpaymentsrequest.CreateInvoiceRequest, userId uint) (*nowpaymentsresponse.CreateInvoiceResponse, error)
	GetInvoiceDataService(page, limit int, isAscendente bool, userId uint) ([]nowpaymentsmodel.NowPaymentCreateInvoice, error)
	GetPaymentStatus(paymentId string) (*nowpaymentsresponse.GetPaymentStatusResponse, error)
	GetInvoiceStatus(invoiceId string) (*nowpaymentsresponse.GetPaymensDataStatusResponse, error)
}
