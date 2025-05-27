package nowpaymentsresponse

import (
	nowpaymentlibs "github.com/jSierraB3991/now_payment_api/domain/now_payment_libs"
)

type InvoiceResponse struct {
	CreateInvoiceId     uint                               `json:"id"`
	NowPaymentId        string                             `json:"now_payment_id"`
	PriceAmount         string                             `json:"price_amount"`
	PriceCurrency       string                             `json:"price_currency"`
	PayCurrency         string                             `json:"pay_currency"`
	Status              nowpaymentlibs.CreateInvoiceStatus `json:"status_pay"`
	NowPaymentPaymentId uint                               `json:"now_payment_payment_id"`
}
