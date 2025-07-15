package nowpaymentsservice

import (
	"context"

	eliotlibs "github.com/jSierraB3991/jsierra-libs"
	nowpaymentsmapper "github.com/jSierraB3991/now_payment_api/domain/now_payments_mapper"
	nowpaymentsresponse "github.com/jSierraB3991/now_payment_api/infrastructure/now_payments_response"
)

func (s *NowPaymentService) GetInvoiceDataService(ctx context.Context, page *eliotlibs.Paggination, params []eliotlibs.PagginationParam) (*nowpaymentsresponse.InvoiceDataResponsePagg, error) {

	data, err := s.repository.GetInvoicePagination(ctx, page, params)
	if err != nil {
		return nil, err
	}

	return &nowpaymentsresponse.InvoiceDataResponsePagg{
		Limit:      page.Limit,
		Page:       page.Page,
		TotalRows:  page.TotalRows,
		TotalPages: page.TotalPages,
		Sort:       page.Sort,
		Data:       nowpaymentsmapper.GetInvoiceDataMapper(data),
	}, nil
}
