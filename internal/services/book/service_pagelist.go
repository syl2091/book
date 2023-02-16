package book

import (
	"book/internal/pkg/core"
	"book/internal/repository/mysql"
	"book/internal/repository/mysql/book_info"
)

type SearchData struct {
	Page     int
	PageSize int
}

func (s *service) PageList(ctx core.Context, data *SearchData) (list []*book_info.BookInfo, err error) {
	page := data.Page
	if page == 0 {
		page = 1
	}
	pageSize := data.PageSize
	if pageSize == 0 {
		pageSize = 10
	}
	offset := (page - 1) * pageSize
	qb := book_info.NewQueryBuilder()
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
