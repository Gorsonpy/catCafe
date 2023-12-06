package pack

import "github.com/Gorsonpy/catCafe/biz/model/order"

func PackOrderBase(resp *order.BaseResponse, code int64, msg string) {
	resp.Code = code
	resp.Msg = msg
}

func PackQueryOrder(resp *order.QueryOrderResp, code int64, msg string, list []*order.OrderModel) {
	resp.Base = order.NewBaseResponse()
	resp.Base.Code = code
	resp.Base.Msg = msg
	resp.Data = make([]*order.OrderModel, 0)
	resp.Data = list
}

func PackCreateOrder(resp *order.CreateOrderResp, code int64, msg string, id int64) {
	resp.Base = order.NewBaseResponse()
	resp.Base.Code = code
	resp.Base.Msg = msg
	resp.Data = make(map[string]int64)
	resp.Data["orderId"] = id
}
