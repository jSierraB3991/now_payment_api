package nowpaymentsrequest

type CreatePaymentRequest struct {
	AmountPrice      float32 `json:"price_amount"`
	PriceCurrency    string  `json:"price_currency"`
	PayCurrency      string  `json:"pay_currency"`
	IpnCallbackURL   string  `json:"ipn_callback_url"`
	OrderID          string  `json:"order_id"`
	OrderDescription string  `json:"order_description"`
}
