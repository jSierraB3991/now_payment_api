package nowpaymentsservice

import (
	"context"

	nowpaymentsresponse "github.com/jSierraB3991/now_payment_api/infrastructure/now_payments_response"
)

func (s *NowPaymentService) GetInvoiceByOrderIdService(ctx context.Context, orderId string) (*nowpaymentsresponse.InvoiceResponse, error) {
	invoice, err := s.repository.GetInvoiceByOrderId(ctx, orderId)
	if err != nil {
		return nil, err
	}

	response := nowpaymentsresponse.InvoiceResponse{
		CreateInvoiceId:     invoice.CreateInvoiceId,
		NowPaymentId:        invoice.NowPaymentId,
		PriceAmount:         invoice.PriceAmount,
		PriceCurrency:       invoice.PriceCurrency,
		PayCurrency:         invoice.PayCurrency,
		Status:              invoice.Status,
		NowPaymentPaymentId: invoice.NowPaymentPaymentId,
	}

	return &response, nil
}
