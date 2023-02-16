package order

import (
	"book/internal/pkg/core"
	"book/internal/repository/mysql"
	"book/internal/repository/mysql/order"
)

func (s *service) Cancel(c core.Context, id int32) (err error) {
	data := map[string]interface{}{
		"status":       2,
		"updated_user": c.SessionUserInfo().UserName,
	}
	qb := order.NewQueryBuilder()
	qb.WhereId(mysql.EqualPredicate, id)
	err = qb.Updates(s.db.GetDbW().WithContext(c.RequestContext()), data)
	if err != nil {
		return err
	}
	return
}
