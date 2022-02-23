package helper

import "orderapi/model"

func ToOrderResponse(o *model.Order) *model.OrderCreatedResp {
	return &model.OrderCreatedResp{
		OrderId:   o.Id,
		CreatedAt: o.CreatedAt,
		Total:     o.Total,
	}
}

// func ToDetailedOrderResponse(o *model.Order) *model.OrderDetailResp {
// 	for i, item := range o.Items {
// 		item.
// 	}
// }
