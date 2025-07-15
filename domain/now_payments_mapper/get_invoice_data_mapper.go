package nowpaymentsmapper

import (
	nowpaymentsmodel "github.com/jSierraB3991/now_payment_api/domain/now_payments_model"
	nowpaymentsresponse "github.com/jSierraB3991/now_payment_api/infrastructure/now_payments_response"
)

func GetInvoiceDataMapper(data []nowpaymentsmodel.NowPaymentCreateInvoice) []nowpaymentsresponse.InvoiceDataResponse {
	var response []nowpaymentsresponse.InvoiceDataResponse
	for _, item := range data {
		response = append(response, nowpaymentsresponse.InvoiceDataResponse{
			Id:                  item.CreateInvoiceId,
			UserId:              item.UserId,
			NowPaymentId:        item.NowPaymentId,
			OrderID:             item.OrderID,
			OrderDescription:    item.OrderDescription,
			PriceAmount:         item.PriceAmount,
			PriceCurrency:       item.PriceCurrency,
			PayCurrency:         item.PayCurrency,
			IpnCallbackURL:      item.IpnCallbackURL,
			InvoiceURL:          item.InvoiceURL,
			SuccessURL:          item.SuccessURL,
			CancelURL:           item.CancelURL,
			CreatedAtMy:         item.CreatedAtMy,
			UpdatedAtMy:         item.UpdatedAtMy,
			Status:              item.Status,
			NowPaymentPaymentId: item.NowPaymentPaymentId,
		})
	}
	return response
}
