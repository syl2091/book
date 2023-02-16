package book

import (
	"book/internal/pkg/core"
	"book/internal/repository/mysql"
	"book/internal/repository/mysql/book_info"
	"book/internal/repository/redis"
)

var _ Service = (*service)(nil)

type Service interface {
	i()
	Create(ctx core.Context, oderData *CreateBookData) (id int32, err error)
	Detail(ctx core.Context, id int32) (book *book_info.BookInfo, err error)
	PageList(ctx core.Context, data *SearchData) (list []*book_info.BookInfo, err error)
	PageListCount(ctx core.Context, data *SearchData) (total int64, err error)
	Delete(ctx core.Context, id int32) (err error)
	Update(ctx core.Context, id *UpdateData) (err error)
}

type service struct {
	db    mysql.Repo
	cache redis.Repo
}

func New(db mysql.Repo, cache redis.Repo) Service {
	return &service{
		db:    db,
		cache: cache,
	}
}
func (s *service) i() {
}
