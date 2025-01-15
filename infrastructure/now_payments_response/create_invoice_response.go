package nowpaymentsresponse

import "time"

type CreateInvoiceResponse struct {
	ID               string    `json:"id"`
	OrderID          string    `json:"order_id"`
	OrderDescription string    `json:"order_description"`
	PriceAmount      string    `json:"price_amount"`
	PriceCurrency    string    `json:"price_currency"`
	PayCurrency      string    `json:"pay_currency"`
	IpnCallbackURL   string    `json:"ipn_callback_url"`
	InvoiceURL       string    `json:"invoice_url"`
	SuccessURL       string    `json:"success_url"`
	CancelURL        string    `json:"cancel_url"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}
