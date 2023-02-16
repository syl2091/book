package book

import (
	"book/internal/pkg/core"
	"book/internal/repository/mysql"
	"book/internal/repository/mysql/book_info"
)

func (s *service) Detail(ctx core.Context, id int32) (book *book_info.BookInfo, err error) {
	qb := book_info.NewQueryBuilder()
	qb.WhereId(mysql.EqualPredicate, id)
	book, err = qb.QueryOne(s.db.GetDbW().WithContext(ctx.RequestContext()))
	if err != nil {
		return nil, err
	}
	return
}
