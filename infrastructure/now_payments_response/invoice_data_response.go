package nowpaymentsresponse

import (
	"time"

	nowpaymentlibs "github.com/jSierraB3991/now_payment_api/domain/now_payment_libs"
)

type InvoiceDataResponse struct {
	Id                  uint                               `json:"id"`
	UserId              uint                               `json:"user_id"`
	NowPaymentId        string                             `json:"now_payment_id"`
	OrderID             string                             `json:"order_id"`
	OrderDescription    string                             `json:"order_description"`
	PriceAmount         string                             `json:"price_amount"`
	PriceCurrency       string                             `json:"price_currency"`
	PayCurrency         string                             `json:"pay_currency"`
	IpnCallbackURL      string                             `json:"ipn_callback_url"`
	InvoiceURL          string                             `json:"invoice_url"`
	SuccessURL          string                             `json:"success_url"`
	CancelURL           string                             `json:"cancel_url"`
	CreatedAtMy         time.Time                          `json:"created_at_my"`
	UpdatedAtMy         time.Time                          `json:"updated_at_my"`
	Status              nowpaymentlibs.CreateInvoiceStatus `json:"status"`
	NowPaymentPaymentId uint                               `json:"now_payment_payment_id"`
}

type InvoiceDataResponsePagg struct {
	Limit      int                   `json:"limit"`
	Page       int                   `json:"page"`
	TotalRows  int64                 `json:"total_rows"`
	TotalPages int                   `json:"total_pages"`
	Sort       string                `json:"-"`
	Data       []InvoiceDataResponse `json:"data"`
}
