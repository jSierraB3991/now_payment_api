package nowpaymentsserviceinterface

import (
	"context"

	nowpaymentsmodel "github.com/jSierraB3991/now_payment_api/domain/now_payments_model"
	nowpaymentsrequest "github.com/jSierraB3991/now_payment_api/infrastructure/now_payments_request"
	nowpaymentsresponse "github.com/jSierraB3991/now_payment_api/infrastructure/now_payments_response"
)

type NowPaymentServiceInterface interface {
	CreatePayment(ctx context.Context, req nowpaymentsrequest.CreatePaymentRequest, userId uint) (*nowpaymentsresponse.CreatePaymentResponse, error)
	CreateInvoice(ctx context.Context, req nowpaymentsrequest.CreateInvoiceRequest, userId uint, payCurrency string) (*nowpaymentsresponse.CreateInvoiceResponse, error)
	GetInvoiceDataService(ctx context.Context, page, limit int, isAscendente bool, userId uint) ([]nowpaymentsmodel.NowPaymentCreateInvoice, error)
	GetPaymentStatus(ctx context.Context, paymentId string) (*nowpaymentsresponse.GetPaymentStatusResponse, error)
	GetInvoiceStatus(ctx context.Context, invoiceId string) (*nowpaymentsresponse.GetPaymentStatusResponse, error)
	GetInvoiceByOrderIdService(ctx context.Context, orderId string) (*nowpaymentsresponse.InvoiceResponse, error)
}
