package nowpaymentsmapper

import (
	nowpaymentsmodel "github.com/jSierraB3991/now_payment_api/domain/now_payments_model"
	nowpaymentsresponse "github.com/jSierraB3991/now_payment_api/infrastructure/now_payments_response"
)

func GetCreateInvoiceByResponse(data nowpaymentsresponse.CreateInvoiceResponse, userId, nowPaymentPaymentId uint) *nowpaymentsmodel.NowPaymentCreateInvoice {
	return &nowpaymentsmodel.NowPaymentCreateInvoice{
		NowPaymentId: data.ID,
		OrderID:      data.OrderID,
		UserId:       userId,

		OrderDescription:    data.OrderDescription,
		PriceAmount:         data.PriceAmount,
		PriceCurrency:       data.PriceCurrency,
		PayCurrency:         data.PayCurrency,
		IpnCallbackURL:      data.IpnCallbackURL,
		InvoiceURL:          data.InvoiceURL,
		SuccessURL:          data.SuccessURL,
		CancelURL:           data.CancelURL,
		CreatedAtMy:         data.CreatedAt,
		UpdatedAtMy:         data.UpdatedAt,
		NowPaymentPaymentId: nowPaymentPaymentId,
	}
}
