package nowpaymentsmapper

import (
	"strings"

	nowpaymentsmodel "github.com/jSierraB3991/now_payment_api/domain/now_payments_model"
)

func PaginationReqToModel(limit, page int, sort, propSort string, dataForWhere interface{}) *nowpaymentsmodel.Paggination {
	return &nowpaymentsmodel.Paggination{
		Limit: limit,
		Page:  page,
		Sort:  GetSort(sort, propSort),
		Data:  dataForWhere,
	}
}

func GetSort(sort, propSort string) string {
	if strings.ToLower(sort) != "asc" && strings.ToLower(sort) != "desc" {
		sort = "desc"
	}

	if propSort == "" {
		propSort = "Id"
	}

	return propSort + " " + sort
}
