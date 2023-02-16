package order

import (
	"book/configs"
	"book/internal/pkg/core"
	"book/internal/repository/mysql"
	"book/internal/repository/redis"
	"book/internal/services/order"
	"book/pkg/hash"

	"go.uber.org/zap"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// Create 创建订单
	// @Tags API.order
	// @Router /api/order/create [post]
	Create() core.HandlerFunc

	// Cancel 取消订单
	// @Tags API.order
	// @Router /api/order/cancel [post]
	Cancel() core.HandlerFunc

	// Detail 订单详情
	// @Tags API.order
	// @Router /api/order/{id} [get]
	Detail() core.HandlerFunc

	List() core.HandlerFunc
}

type handler struct {
	logger       *zap.Logger
	cache        redis.Repo
	hashids      hash.Hash
	orderService order.Service
}

func New(logger *zap.Logger, db mysql.Repo, cache redis.Repo) Handler {
	return &handler{
		logger:       logger,
		cache:        cache,
		hashids:      hash.New(configs.Get().HashIds.Secret, configs.Get().HashIds.Length),
		orderService: order.New(db, cache),
	}
}

func (h *handler) i() {}
