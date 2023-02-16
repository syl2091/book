package book

import (
	"book/internal/code"
	"book/internal/pkg/core"
	"net/http"
	"time"
)

type detailRequest struct {
	Id int32 `form:"id"`
}

type detailResponse struct {
	Name         string    `json:"name" `
	Author       string    `json:"author"`
	Publish      string    `json:"publish"`
	ISBN         string    `json:"ISBN"`
	Introduction string    `json:"introduction"`
	Language     string    `json:"language"`
	Price        float64   `json:"price"`
	Pubdate      time.Time `json:"pubdate"`
	ClassId      int32     `json:"classId"`
	Pressmark    int32     `json:"pressmark"`
}

func (h *handler) Detail() core.HandlerFunc {
	return func(c core.Context) {
		req := new(detailRequest)
		res := new(detailResponse)
		if err := c.ShouldBindForm(req); err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithError(err),
			)
			return
		}

		detail, err := h.bookService.Detail(c, req.Id)
		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.BookDetailError,
				code.Text(code.BookDetailError)).WithError(err),
			)
			return
		}
		res.Name = detail.Name
		res.Author = detail.Author
		res.Publish = detail.Publish
		res.ISBN = detail.ISBN
		res.Introduction = detail.Introduction
		res.Price = detail.Price
		res.Language = detail.Language
		res.Pubdate = detail.Pubdate
		res.ClassId = detail.ClassId
		res.Pressmark = detail.Pressmark
		c.Payload(res)

	}
}
