package mysql

import (
	"time"

	"github.com/Gorsonpy/catCafe/biz/model/order"
)

type Order struct {
	OrderId       int64 `gorm:"primary_key"`
	CustomerId    int64
	CatId         int64
	OrderDetails  string
	TotalAmount   float64
	PaymentStatus bool
	PaymentMethod string
	OrderTime     time.Time
}

func OrderToModel(o *Order) *order.OrderModel {
	model := order.NewOrderModel()
	model.OrderId = o.OrderId
	model.CustomerId = o.CustomerId
	model.OrderDetails = o.OrderDetails
	model.TotalAmount = o.TotalAmount
	model.PaymentStatus = o.PaymentStatus
	model.PaymentMethod = o.PaymentMethod
	model.OrderTime = o.OrderTime.Format(time.DateTime)
	return model
}
func QueryOrders(cId int64, status bool, limit int64) ([]*Order, error) {
	list := make([]*Order, 0)
	err := DB.Table("orders").Limit(int(limit)).
		Where("customer_id = ? and payment_status = ?", cId, status).
		Find(&list).Error
	return list, err
}

func CreateOrder(o *Order) error {
	return DB.Table("orders").Create(&o).Error
}

func UpdateOrder(id int64, orderDeltails string, totalAmount float64) error {
	return DB.Table("orders").Where("order_id = ?", id).
		Updates(map[string]interface{}{"order_details": orderDeltails, "total_amount": totalAmount}).Error
}

func UpdateOrderSt(id int64, status bool) error {
	return DB.Table("orders").Where("order_id = ?", id).
		Update("payment_status", status).Error
}

func DelOrders(id []int64) error {
	return DB.Table("orders").Delete(&Order{}, id).Error
}
