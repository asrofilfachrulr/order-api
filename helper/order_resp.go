package helper

import "orderapi/model"

func ToOrderResponse(o *model.Order) *model.OrderResp {
	return &model.OrderResp{
		OrderId:   o.Id,
		CreatedAt: o.CreatedAt,
		Total:     o.Total,
	}
}
