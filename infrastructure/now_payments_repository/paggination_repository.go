package nowpaymentsrepository

import (
	"math"

	nowpaymentsmodel "github.com/jSierraB3991/now_payment_api/domain/now_payments_model"
	"gorm.io/gorm"
)

func (repo *Repository) Paginate(value interface{}, page *nowpaymentsmodel.Paggination, args []nowpaymentsmodel.PagginationParam, preloads []nowpaymentsmodel.PreloadParams) error {
	return repo.db.Scopes(repo.paginate_with_param(value, page, args, preloads)).Find(&value).Error
}

func (repo *Repository) paginate_with_param(value interface{}, page *nowpaymentsmodel.Paggination, args []nowpaymentsmodel.PagginationParam, preloads []nowpaymentsmodel.PreloadParams) func(db *gorm.DB) *gorm.DB {
	var totalRows int64
	accountData := repo.db.Model(value)
	if len(args) > 0 {
		for _, arg := range args {
			accountData.Where(arg.Where, arg.Data)
		}
	}
	accountData.Count(&totalRows)

	page.TotalRows = totalRows
	totalPages := int(math.Ceil(float64(totalRows) / float64(page.Limit)))
	page.TotalPages = totalPages

	return func(db *gorm.DB) *gorm.DB {
		data := db.Offset(page.GetOffset()).Limit(page.GetLimit()).Order(page.GetSort())

		if len(preloads) > 0 {
			for _, arg := range preloads {
				if arg.PagginationParam.Where == "" {
					data.Preload(arg.Preload)
				} else {
					data.Preload(arg.Preload, arg.PagginationParam.Where, arg.PagginationParam.Data)
				}
			}
		}

		if len(args) > 0 {
			for _, arg := range args {
				data.Where(arg.Where, arg.Data)
			}
		}
		return data
	}
}
