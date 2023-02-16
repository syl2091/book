package order

import (
	"net/http"
	"time"

	"book/internal/code"
	"book/internal/pkg/core"
)

type detailRequest struct {
	Id int32 `uri:"id" `
}

type detailResponse struct {
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

// Detail 取消订单
// @Summary 取消订单
// @Description 取消订单
// @Tags API.order
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body detailRequest true "请求信息"
// @Success 200 {object} detailResponse
// @Failure 400 {object} code.Failure
// @Router /api/order/{id} [get]
func (h *handler) Detail() core.HandlerFunc {
	return func(ctx core.Context) {
		req := new(detailRequest)
		res := new(detailResponse)

		if err := ctx.ShouldBindURI(req); err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithError(err))
			return
		}

		info, err := h.orderService.Detail(ctx, req.Id)
		if err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.OrderDetailError,
				code.Text(code.OrderDetailError)).WithError(err))
			return
		}

		res.Id = info.Id
		res.OrderNo = info.OrderNo
		res.OrderFee = info.OrderFee
		res.CreatedAt = info.CreatedAt
		res.CreatedUser = info.CreatedUser
		res.Status = info.Status
		res.IsDeleted = info.IsDeleted
		res.UpdatedAt = info.UpdatedAt
		res.UpdatedUser = info.UpdatedUser
		ctx.Payload(res)
	}
}
