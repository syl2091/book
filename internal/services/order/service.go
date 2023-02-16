package order

import (
	"book/internal/pkg/core"
	"book/internal/repository/mysql"
	"book/internal/repository/mysql/order"
	"book/internal/repository/redis"
)

var _ Service = (*service)(nil)

type Service interface {
	i()
	Create(ctx core.Context, oderData *CreateOrderData) (id int32, err error)
	Cancel(ctx core.Context, id int32) (err error)
	Detail(ctx core.Context, id int32) (order *order.Order, err error)
	PageList(ctx core.Context, data *SearchData) (list []*order.Order, err error)
	PageListCount(ctx core.Context, data *SearchData) (total int64, err error)
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
