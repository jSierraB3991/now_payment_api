package nowpaymentsservice

import (
	"context"

	eliotlibs "github.com/jSierraB3991/jsierra-libs"
	nowpaymentsmapper "github.com/jSierraB3991/now_payment_api/domain/now_payments_mapper"
	nowpaymentsmodel "github.com/jSierraB3991/now_payment_api/domain/now_payments_model"
)

func (s *NowPaymentService) GetInvoiceDataService(ctx context.Context, page, limit int, isAscendente bool, params []eliotlibs.PagginationParam) ([]nowpaymentsmodel.NowPaymentCreateInvoice, error) {

	order := "desc"
	if isAscendente {
		order = "asc"
	}
	paggination := nowpaymentsmapper.PaginationReqToModel(limit, page, order, "order_id", nil)
	data, err := s.repository.GetInvoicePagination(ctx, paggination, params)
	if err != nil {
		return nil, err
	}
	return data, nil
}
