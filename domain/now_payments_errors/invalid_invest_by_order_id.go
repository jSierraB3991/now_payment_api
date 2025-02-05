package nowpaymentserrors

type InvalidInvestOrderIdError struct {
	OrderId string
}

func (err InvalidInvestOrderIdError) Error() string {
	return "INVALID_INVEST_BY_ORDER_ID_" + err.OrderId
}
