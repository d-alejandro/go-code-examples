package types

import "database/sql/driver"

type OrderStatus string

const (
	Canceled   OrderStatus = "canceled"
	Completed  OrderStatus = "completed"
	Paid       OrderStatus = "paid"
	Prepayment OrderStatus = "prepayment"
	Waiting    OrderStatus = "waiting"
)

func (orderStatus *OrderStatus) Scan(value interface{}) error {
	*orderStatus = OrderStatus(value.(string))
	return nil
}

func (orderStatus OrderStatus) Value() (driver.Value, error) {
	return string(orderStatus), nil
}
