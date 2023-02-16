package book

import (
	"book/internal/code"
	"book/internal/pkg/core"
	"book/internal/services/book"
	"github.com/spf13/cast"
	"net/http"
	"time"
)

type listRequest struct {
	Page     int `form:"page"`      // 第几页
	PageSize int `form:"page_size"` // 每页显示条数
}

type listData struct {
	Id           int32     `json:"id"` // 主键
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
	State        int32     `json:"state"`
	IsDeleted    int32     `json:"isDeleted"`
	CreatedAt    time.Time `json:"createdAt"`
	CreatedUser  string    `json:"createdUser"`
	UpdatedAt    time.Time `json:"updatedAt"`
	UpdatedUser  string    `json:"updatedUser"`
}

type listResponse struct {
	List       []listData `json:"list"`
	Pagination struct {
		Total        int `json:"total"`
		CurrentPage  int `json:"current_page"`
		PerPageCount int `json:"per_page_count"`
	} `json:"pagination"`
}

func (h *handler) List() core.HandlerFunc {
	return func(c core.Context) {
		req := new(listRequest)
		res := new(listResponse)

		if err := c.ShouldBindForm(req); err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithError(err),
			)
			return
		}
		page := req.Page
		if page == 0 {
			page = 1
		}

		pageSize := req.PageSize
		if pageSize == 0 {
			pageSize = 10
		}
		serchData := new(book.SearchData)
		serchData.PageSize = pageSize
		serchData.Page = page

		list, err := h.bookService.PageList(c, serchData)
		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.OrderListError,
				code.Text(code.AdminListError)).WithError(err),
			)
			return
		}

		count, err := h.bookService.PageListCount(c, serchData)
		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.OrderListError,
				code.Text(code.OrderListError)).WithError(err),
			)
			return
		}

		res.Pagination.Total = cast.ToInt(count)
		res.Pagination.PerPageCount = pageSize
		res.Pagination.CurrentPage = page
		res.List = make([]listData, len(list))

		for k, v := range list {
			data := listData{
				Id:           v.Id,
				Name:         v.Name,
				Author:       v.Author,
				State:        v.State,
				Publish:      v.Publish,
				ISBN:         v.ISBN,
				Introduction: v.Introduction,
				Language:     v.Language,
				Price:        v.Price,
				Pubdate:      v.Pubdate,
				ClassId:      v.ClassId,
				Pressmark:    v.Pressmark,
				IsDeleted:    v.IsDeleted,
				CreatedAt:    v.CreatedAt,
				CreatedUser:  v.CreatedUser,
				UpdatedAt:    v.UpdatedAt,
				UpdatedUser:  v.UpdatedUser,
			}
			res.List[k] = data
		}
		c.Payload(res)
	}
}
