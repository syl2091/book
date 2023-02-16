package book

import (
	"book/internal/code"
	"book/internal/pkg/core"
	"net/http"
)

type deleteRequest struct {
	Id int32 `form:"id"`
}

type deleteResponse struct {
	Id int32 `json:"id"`
}

func (h *handler) Delete() core.HandlerFunc {
	return func(c core.Context) {
		req := new(deleteRequest)
		res := new(deleteResponse)
		if err := c.ShouldBindForm(req); err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithError(err),
			)
			return
		}
		err := h.bookService.Delete(c, req.Id)
		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.BookDeleteError,
				code.Text(code.BookDeleteError)).WithError(err),
			)
			return
		}
		res.Id = req.Id
		c.Payload(res)
	}
}
