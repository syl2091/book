package book

import (
	"book/internal/code"
	"book/internal/pkg/core"
	"book/internal/repository/mysql"
	"book/internal/repository/redis"
	"fmt"
	"go.uber.org/zap"
	"net/http"
)

type detailReq struct {
	id int32 `uri:"id"`
}

type handler struct {
	db     mysql.Repo
	logger *zap.Logger
	cache  redis.Repo
}

func New(logger *zap.Logger, db mysql.Repo, cache redis.Repo) *handler {
	return &handler{
		logger: logger,
		cache:  cache,
		db:     db,
	}
}

func (h *handler) List() core.HandlerFunc {
	return func(ctx core.Context) {
		ctx.HTML("book_list", nil)
	}
}
func (h *handler) Create() core.HandlerFunc {
	return func(ctx core.Context) {
		ctx.HTML("book_add", nil)
	}
}
func (h *handler) Detail() core.HandlerFunc {
	return func(ctx core.Context) {
		type editRequest struct {
			Id int32 `uri:"id"`
		}

		type editResponse struct {
			Id int32 `json:"id"`
		}
		req := new(editRequest)
		res := new(editResponse)
		if err := ctx.ShouldBindURI(req); err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithError(err),
			)
			return
		}
		fmt.Println(req.Id)
		res.Id = req.Id
		ctx.HTML("book_detail", res)
	}
}
