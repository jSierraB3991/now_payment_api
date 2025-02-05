package nowpaymentsresponse

import "time"

type GetPaymentStatusResponse struct {
	PaymentId        uint      `json:"payment_id"`
	InvoiceId        *string   `json:"invoice_id"`
	PaymentStatus    string    `json:"payment_status"`
	PayAddress       string    `json:"pay_address"`
	PayInExtraId     *string   `json:"payin_extra_id"`
	PriceAmount      float32   `json:"price_amount"`
	PriceCurrency    string    `json:"price_currency"`
	PayAmount        float32   `json:"pay_amount"`
	ActualPaid       uint      `json:"actually_paid"`
	PayCurrency      string    `json:"pay_currency"`
	OrderId          *string   `json:"order_id"`
	OrderDescription *string   `json:"order_description"`
	PurchaseId       uint      `json:"purchase_id"`
	OutComeAmount    float32   `json:"outcome_amount"`
	OutComeCurrency  string    `json:"outcome_currency"`
	PayoutHash       string    `json:"payout_hash"`
	PayinHash        string    `json:"payin_hash"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	BurningPercent   string    `json:"burning_percent"`
	Type             string    `json:"type"`
	PaymentExtraIds  []int64   `json:"payment_extra_ids"`
}

type GetPaymensDataStatusResponse struct {
	Data       []GetPaymentStatusResponse `json:"data"`
	Limit      int                        `json:"limit"`
	Page       uint                       `json:"page"`
	PagesCount uint                       `json:"pagesCount"`
	Total      uint                       `json:"total"`
}
