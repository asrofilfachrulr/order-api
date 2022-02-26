package helper

import (
	"orderapi/model"
	"time"
)

func ToOrderCreatedResponse(o *model.Order) *model.OrderCreatedResp {
	return &model.OrderCreatedResp{
		OrderId:   o.Id,
		CreatedAt: o.CreatedAt,
		Total:     o.Total,
	}
}

func ToOrdersBriefResponse(os []model.Order) []model.OrderBrief {
	obs := []model.OrderBrief{}

	for _, order := range os {
		obs = append(obs, model.OrderBrief{
			Id:        order.Id,
			Status:    order.Status,
			UpdatedAt: order.UpdatedAt.Format(time.RFC1123),
			Total:     order.Total,
		})
	}
	return obs
}
