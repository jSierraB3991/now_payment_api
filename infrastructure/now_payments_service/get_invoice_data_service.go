package nowpaymentsservice

import (
	nowpaymentsmapper "github.com/jSierraB3991/now_payment_api/domain/now_payments_mapper"
	nowpaymentsmodel "github.com/jSierraB3991/now_payment_api/domain/now_payments_model"
)

func (s *NowPaymentService) GetInvoiceDataService(page, limit int, isAscendente bool) ([]nowpaymentsmodel.NowPaymentCreateInvoice, error) {

	order := "desc"
	if isAscendente {
		order = "asc"
	}
	paggination := nowpaymentsmapper.PaginationReqToModel(limit, page, order, "order_id", nil)
	data, err := s.repository.GetInvoicePagination(paggination)
	if err != nil {
		return nil, err
	}
	return data, nil
}
