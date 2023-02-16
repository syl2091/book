package book

import (
	"book/internal/code"
	"book/internal/pkg/core"
	"book/internal/services/book"
	"net/http"
	"time"
)

type updateRequest struct {
	Id           int32   `form:"id"`
	Name         string  `form:"name" `
	Author       string  `form:"author"`
	Publish      string  `form:"publish"`
	ISBN         string  `form:"ISBN"`
	Introduction string  `form:"introduction"`
	Language     string  `form:"language"`
	Price        float64 `form:"price"`
	Pubdate      string  `form:"pubdate"`
	ClassId      int32   `form:"classId"`
	Pressmark    int32   `form:"pressmark"`
}

type updateResponse struct {
	Id int32 `json:"id"`
}

func (h *handler) Update() core.HandlerFunc {
	return func(c core.Context) {
		req := new(updateRequest)
		res := new(updateResponse)
		if err := c.ShouldBindForm(req); err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithError(err),
			)
			return
		}
		data := new(book.UpdateData)
		data.Id = req.Id
		data.Name = req.Name
		data.Author = req.Author
		data.Publish = req.Publish
		data.ISBN = req.ISBN
		data.Introduction = req.Introduction
		data.Price = req.Price
		data.Language = req.Language
		time, _ := time.ParseInLocation("2006-01-02 15:04:05", req.Pubdate, time.Local)
		data.Pubdate = time
		data.ClassId = req.ClassId
		data.Pressmark = req.Pressmark
		err := h.bookService.Update(c, data)
		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.BookUpdateError,
				code.Text(code.BookUpdateError)).WithError(err),
			)
			return
		}
		res.Id = req.Id
		c.Payload(res)
	}
}
