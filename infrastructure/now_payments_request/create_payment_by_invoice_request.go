package nowpaymentsrequest

type CreatePaymentByInvoiceRequest struct {
	Iid         string `json:"iid"`
	PayCurrency string `json:"pay_currency"`
}
