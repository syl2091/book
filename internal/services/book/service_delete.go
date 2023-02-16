package book

import (
	"book/internal/pkg/core"
	"book/internal/repository/mysql"
	"book/internal/repository/mysql/book_info"
)

func (s *service) Delete(ctx core.Context, id int32) (err error) {
	data := map[string]interface{}{
		"is_deleted":   1,
		"updated_user": ctx.SessionUserInfo().UserName,
	}
	qb := book_info.NewQueryBuilder()
	qb.WhereId(mysql.EqualPredicate, id)
	err = qb.Updates(s.db.GetDbR().WithContext(ctx.RequestContext()), data)
	if err != nil {
		return err
	}
	return
}
