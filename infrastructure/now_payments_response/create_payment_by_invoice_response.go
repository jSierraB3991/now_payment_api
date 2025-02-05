package nowpaymentsresponse

type CreatePaymentByInvoiceResponse struct {
	PaymentId     string `json:"payment_id"`
	PaymentStatus string `json:"payment_status"`
}
