package nowpaymentsrepository

import (
	"context"
	"math"

	jsierralibs "github.com/jSierraB3991/jsierra-libs"
	"gorm.io/gorm"
)

func (repo *Repository) paginate_with_param(ctx context.Context, value interface{}, page *jsierralibs.Paggination,
	args []jsierralibs.PagginationParam, preloads []jsierralibs.PreloadParams) func(db *gorm.DB) *gorm.DB {

	db, _ := repo.GetDb(ctx)
	var totalRows int64
	accountData := db.Model(value)
	if len(preloads) > 0 {
		for _, arg := range preloads {
			if arg.PagginationParam.Where == "" {
				accountData.Preload(arg.Preload)
			} else {
				accountData.Preload(arg.Preload, arg.PagginationParam.Where, arg.PagginationParam.Data)
			}
		}
	}
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
