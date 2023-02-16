package book

import (
	"book/configs"
	"book/internal/pkg/core"
	"book/internal/repository/mysql"
	"book/internal/repository/redis"
	"book/internal/services/book"
	"book/pkg/hash"

	"go.uber.org/zap"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// Create 创建书籍
	// @Tags API.order
	// @Router /api/order/create [post]
	Create() core.HandlerFunc

	// Detail 删除书籍
	// @Tags API.order
	// @Router /api/order/delete [get]
	Delete() core.HandlerFunc

	Update() core.HandlerFunc

	List() core.HandlerFunc

	Detail() core.HandlerFunc
}

type handler struct {
	logger      *zap.Logger
	db          mysql.Repo
	cache       redis.Repo
	hashids     hash.Hash
	bookService book.Service
}

func New(logger *zap.Logger, db mysql.Repo, cache redis.Repo) Handler {
	return &handler{
		logger:      logger,
		cache:       cache,
		hashids:     hash.New(configs.Get().HashIds.Secret, configs.Get().HashIds.Length),
		bookService: book.New(db, cache),
	}
}

func (h *handler) i() {}
