package order

import (
	"net/http"

	"book/internal/code"
	"book/internal/pkg/core"
)

type cancelRequest struct {
	Id int32 `json:"id" binding:"required"`
}

type cancelResponse struct {
	msg string `json:"msg"`
}

// Cancel 取消订单
// @Summary 取消订单
// @Description 取消订单
// @Tags API.order
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body cancelRequest true "请求信息"
// @Success 200 {object} cancelResponse
// @Failure 400 {object} code.Failure
// @Router /api/order/cancel [post]
func (h *handler) Cancel() core.HandlerFunc {
	return func(ctx core.Context) {
		req := new(cancelRequest)
		res := new(cancelResponse)
		if err := ctx.ShouldBindJSON(req); err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithError(err),
			)
			return
		}
		err := h.orderService.Cancel(ctx, req.Id)
		if err != nil {
			return
		}
		res.msg = "取消订单成功!"
		ctx.Payload(res)
	}
}
