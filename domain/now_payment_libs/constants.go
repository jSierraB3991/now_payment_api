package nowpaymentlibs

const (
	CREATE_PAYMENT_URL        = "/payment"
	CREATE_INVOICE_URL        = "/invoice"
	CREATE_PAYMENT_BY_INVOICE = "/invoice-payment"
	GET_PAYMENT_DATA_URL      = "/payment/"
	AUTH_URL                  = "/auth"

	NOW_PAYMENTS_STATUS_RESULT_WAITING    string = "waiting"
	NOW_PAYMENTS_STATUS_RESULT_CONFIRMING string = "confirming"
	NOW_PAYMENTS_STATUS_RESULT_CONFIRMED  string = "confirmed"
	NOW_PAYMENTS_STATUS_RESULT_FINISHED   string = "finished"
	NOW_PAYMENTS_STATUS_RESULT_FAILED     string = "failed"
	NOW_PAYMENTS_STATUS_RESULT_EXPIRED    string = "expired"
)
