package book

import (
	"book/internal/code"
	"book/internal/pkg/core"
	"book/internal/services/book"
	"net/http"
	"time"
)

type createRequest struct {
	Name         string  `form:"name" `        //
	Author       string  `form:"author"`       //
	Publish      string  `form:"publish"`      //
	ISBN         string  `form:"ISBN"`         //
	Introduction string  `form:"introduction"` //
	Language     string  `form:"language"`     //
	Price        float64 `form:"price"`        //
	Pubdate      string  `form:"pubdate"`
	ClassId      int32   `form:"classId"`   //
	Pressmark    int32   `form:"pressmark"` //
}

type createResponse struct {
	Id int32 `json:"id"`
}

// Create 创建书籍
// @Summary 创建书籍
// @Description 创建书籍
// @Tags API.order
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body createRequest true "请求信息"
// @Success 200 {object} createResponse
// @Failure 400 {object} code.Failure
// @Router /api/order/create [post]
func (h *handler) Create() core.HandlerFunc {
	return func(ctx core.Context) {
		req := new(createRequest)
		res := new(createResponse)
		if err := ctx.ShouldBindForm(req); err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithError(err),
			)
			return
		}
		bookData := new(book.CreateBookData)
		bookData.Name = req.Name
		bookData.Author = req.Author
		bookData.Publish = req.Publish
		bookData.ISBN = req.ISBN
		bookData.Introduction = req.Introduction
		bookData.Language = req.Language
		bookData.Price = req.Price
		time, _ := time.ParseInLocation("2006-01-02 15:04:05", req.Pubdate, time.Local)
		bookData.Pubdate = time
		bookData.ClassId = req.ClassId
		bookData.Pressmark = req.Pressmark

		id, err := h.bookService.Create(ctx, bookData)
		if err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.BookCreateError,
				code.Text(code.BookCreateError)).WithError(err),
			)
			return
		}
		res.Id = id
		ctx.Payload(res)

	}
}
