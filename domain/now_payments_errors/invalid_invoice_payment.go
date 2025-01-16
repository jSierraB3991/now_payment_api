package nowpaymentserrors

type InvalidInvoicePaymentError struct{}

func (InvalidInvoicePaymentError) Error() string {
	return "INVALID_INVOICE_PAYMENT_ERROR"
}
