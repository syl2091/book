package order

import (
	"book/internal/pkg/core"
	"book/internal/repository/mysql"
	"book/internal/repository/mysql/order"
)

func (s *service) PageListCount(ctx core.Context, data *SearchData) (total int64, err error) {
	qb := order.NewQueryBuilder()
	qb = qb.WhereIsDeleted(mysql.EqualPredicate, -1)

	total, err = qb.Count(s.db.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return 0, err
	}
	return
}
