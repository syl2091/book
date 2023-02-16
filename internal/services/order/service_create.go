package order

import (
	"book/internal/pkg/core"
	"book/internal/repository/mysql/order"
)

type CreateOrderData struct {
	OrderNo  string // 订单号
	OrderFee int32  // 订单金额(分)
}

func (s *service) Create(ctx core.Context, orderdata *CreateOrderData) (id int32, err error) {
	model := order.NewModel()
	model.OrderNo = orderdata.OrderNo
	model.OrderFee = orderdata.OrderFee
	model.Status = 1
	model.CreatedUser = ctx.SessionUserInfo().UserName
	model.IsDeleted = -1

	id, err = model.Create(s.db.GetDbW().WithContext(ctx.RequestContext()))
	if err != nil {
		return 0, err
	}
	return
}
