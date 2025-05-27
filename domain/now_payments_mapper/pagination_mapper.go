package nowpaymentsmapper

import (
	"strings"

	jsierralibs "github.com/jSierraB3991/jsierra-libs"
)

func PaginationReqToModel(limit, page int, sort, propSort string, dataForWhere interface{}) *jsierralibs.Paggination {
	return &jsierralibs.Paggination{
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
