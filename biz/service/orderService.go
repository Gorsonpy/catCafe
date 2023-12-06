package service

import (
	"time"

	"github.com/Gorsonpy/catCafe/biz/dal/mysql"
	"github.com/Gorsonpy/catCafe/biz/model/order"
	"github.com/Gorsonpy/catCafe/pkg/errno"
	"gorm.io/gorm"
)

func QueryOrders(cId int64, st bool, limit int64) (int64, string, []*order.OrderModel) {
	list := make([]*order.OrderModel, 0)
	orders, err := mysql.QueryOrders(cId, st, limit)
	if err != gorm.ErrRecordNotFound && err != nil {
		return errno.GetErrorCode, err.Error(), nil
	}
	for _, val := range orders {
		list = append(list, mysql.OrderToModel(val))
	}
	return errno.StatusSuccessCode, errno.SuccessMsg, list
}

func CreateOrder(customerId int64, req *order.CreateOrderReq) (int64, string, int64) {
	o := mysql.Order{
		CatId:         req.CatId,
		CustomerId:    customerId,
		TotalAmount:   req.TotalAmount,
		OrderDetails:  req.OrderDetails,
		PaymentMethod: req.PaymentMethod,
		OrderTime:     time.Now(),
	}
	err := mysql.CreateOrder(&o)
	if err != nil {
		return errno.CreateErrorCode, err.Error(), 0
	}
	return errno.StatusSuccessCode, errno.SuccessMsg, o.OrderId
}

func UpdateOrder(req *order.UpdateOrderReq) (int64, string) {
	err := mysql.UpdateOrder(req.OrderId, req.OrderDetails, req.TotalAmount)
	if err != nil {
		return errno.UpdateErrorCode, err.Error()
	}
	return errno.StatusSuccessCode, errno.SuccessMsg
}

func ConfirmOrder(req *order.ConfirmOrderReq) (int64, string) {
	err := mysql.UpdateOrderSt(req.OrderId, req.GetStatus())
	if err != nil {
		return errno.UpdateErrorCode, err.Error()
	}
	return errno.StatusSuccessCode, errno.SuccessMsg
}

func DelOrder(ids []int64) (int64, string) {
	err := mysql.DelOrders(ids)
	if err != nil {
		return errno.DelErrorCode, err.Error()
	}
	return errno.StatusSuccessCode, errno.SuccessMsg
}
