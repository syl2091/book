package order

import (
	"net/http"
	"time"

	"book/internal/code"
	"book/internal/pkg/core"
	"book/internal/services/order"

	"github.com/spf13/cast"
)

type listRequest struct {
	Page     int `form:"page"`      // 第几页
	PageSize int `form:"page_size"` // 每页显示条数
}

type listData struct {
	Id          int32     `json:"id"`          // 主键
	OrderNo     string    `json:"orderNo"`     // 订单号
	OrderFee    int32     `json:"orderFee"`    // 订单金额(分)
	Status      int32     `json:"status"`      // 订单状态 1:已创建  2:已取消
	IsDeleted   int32     `json:"isDeleted"`   // 是否删除 1:是  -1:否
	CreatedAt   time.Time `json:"createdAt"`   // 创建时间
	CreatedUser string    `json:"createdUser"` // 创建人
	UpdatedAt   time.Time `json:"updatedAt"`   // 更新时间
	UpdatedUser string    `json:"updatedUser"` // 更新人
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
		serchData := new(order.SearchData)
		serchData.PageSize = pageSize
		serchData.Page = page

		list, err := h.orderService.PageList(c, serchData)
		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.OrderListError,
				code.Text(code.AdminListError)).WithError(err),
			)
			return
		}

		count, err := h.orderService.PageListCount(c, serchData)
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
				Id:          v.Id,
				OrderNo:     v.OrderNo,
				OrderFee:    v.OrderFee,
				Status:      v.Status,
				IsDeleted:   v.IsDeleted,
				CreatedAt:   v.CreatedAt,
				CreatedUser: v.CreatedUser,
				UpdatedAt:   v.UpdatedAt,
				UpdatedUser: v.UpdatedUser,
			}
			res.List[k] = data
		}
		c.Payload(res)
	}
}
