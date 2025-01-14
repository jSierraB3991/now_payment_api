package nowpaymentsmapper

import (
	nowpaymentsmodel "github.com/jSierraB3991/now_payment_api/domain/now_payments_model"
	nowpaymentsresponse "github.com/jSierraB3991/now_payment_api/infrastructure/now_payments_response"
)

func GetCreatePaymentByCreatePaymentResponse(data nowpaymentsresponse.CreatePaymentResponse, userId uint) nowpaymentsmodel.NowPaymentCreatePayment {
	return nowpaymentsmodel.NowPaymentCreatePayment{
		PaymentID:              data.PaymentID,
		PaymentStatus:          data.PaymentStatus,
		PayAddress:             data.PayAddress,
		PriceAmount:            data.PriceAmount,
		PriceCurrency:          data.PriceCurrency,
		PayAmount:              data.PayAmount,
		OrderID:                data.OrderID,
		PayCurrency:            data.PayCurrency,
		OrderDescription:       data.OrderDescription,
		IpnCallbackURL:         data.IpnCallbackURL,
		CreatedAt:              data.CreatedAt,
		PurchaseID:             data.PurchaseID,
		AmountReceived:         data.AmountReceived,
		PayinExtraID:           data.PayinExtraID,
		SmartContract:          data.SmartContract,
		Network:                data.Network,
		UpdatedAt:              data.UpdatedAt,
		NetworkPrecision:       data.NetworkPrecision,
		TimeLimit:              data.TimeLimit,
		BurningPercent:         data.BurningPercent,
		ExpirationEstimateDate: data.ExpirationEstimateDate,
		UserId:                 userId,
	}
}
