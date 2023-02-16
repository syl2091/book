package order

import (
	"book/internal/pkg/core"
	"book/internal/repository/mysql"
	"book/internal/repository/mysql/order"
)

type SearchData struct {
	Page     int // 第几页
	PageSize int // 每页显示条数

}

func (s *service) PageList(ctx core.Context, data *SearchData) (list []*order.Order, err error) {

	page := data.Page
	if page == 0 {
		page = 1
	}

	pageSize := data.PageSize
	if pageSize == 0 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	qb := order.NewQueryBuilder()
	qb.WhereIsDeleted(mysql.EqualPredicate, -1)

	list, err = qb.
		Limit(pageSize).
		Offset(offset).
		OrderById(false).
		QueryAll(s.db.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return nil, err
	}

	return
}
