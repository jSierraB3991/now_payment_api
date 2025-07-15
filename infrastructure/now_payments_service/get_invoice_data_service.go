package nowpaymentsservice

import (
	"context"

	eliotlibs "github.com/jSierraB3991/jsierra-libs"
	nowpaymentsmodel "github.com/jSierraB3991/now_payment_api/domain/now_payments_model"
)

func (s *NowPaymentService) GetInvoiceDataService(ctx context.Context, page *eliotlibs.Paggination, params []eliotlibs.PagginationParam) ([]nowpaymentsmodel.NowPaymentCreateInvoice, error) {

	data, err := s.repository.GetInvoicePagination(ctx, page, params)
	if err != nil {
		return nil, err
	}
	return data, nil
}
