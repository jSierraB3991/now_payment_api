package nowpaymentsresponse

import "time"

type CreatePaymentResponse struct {
	PaymentID              string    `json:"payment_id"`
	PaymentStatus          string    `json:"payment_status"`
	PayAddress             string    `json:"pay_address"`
	PriceAmount            float32   `json:"price_amount"`
	PriceCurrency          string    `json:"price_currency"`
	PayAmount              float32   `json:"pay_amount"`
	PayCurrency            string    `json:"pay_currency"`
	OrderID                string    `json:"order_id"`
	OrderDescription       string    `json:"order_description"`
	IpnCallbackURL         string    `json:"ipn_callback_url"`
	CreatedAt              time.Time `json:"created_at"`
	UpdatedAt              time.Time `json:"updated_at"`
	PurchaseID             string    `json:"purchase_id"`
	AmountReceived         float32   `json:"amount_received"`
	PayinExtraID           string    `json:"payin_extra_id"`
	SmartContract          string    `json:"smart_contract"`
	Network                string    `json:"network"`
	NetworkPrecision       int       `json:"network_precision"`
	TimeLimit              string    `json:"time_limit"`
	BurningPercent         string    `json:"burning_percent"`
	ExpirationEstimateDate time.Time `json:"expiration_estimate_date"`
}
