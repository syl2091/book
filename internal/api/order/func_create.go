package order

import (
	"net/http"

	"book/internal/code"
	"book/internal/pkg/core"
	"book/internal/pkg/validation"
	"book/internal/services/order"
)

type createRequest struct {
	OrderNo   string `form:"orderno"  binding:"required"`   // 订单号
	OrderFree int32  `form:"orderfree"  binding:"required"` // 订单金额
}

type createResponse struct {
	Id int32 `json:"id"`
}

// Create 创建订单
// @Summary 创建订单
// @Description 创建订单
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
				validation.Error(err)).WithError(err),
			)
			return
		}

		createOrder := new(order.CreateOrderData)
		createOrder.OrderNo = req.OrderNo
		createOrder.OrderFee = req.OrderFree
		id, err := h.orderService.Create(ctx, createOrder)
		if err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.OrderCreateError,
				code.Text(code.OrderCreateError)).WithError(err),
			)
			return
		}
		res.Id = id
		ctx.Payload(res)
	}
}
