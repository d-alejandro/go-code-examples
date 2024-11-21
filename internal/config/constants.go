package config

import "github.com/d-alejandro/go-code-examples/internal/pkg/models/types"

const (
	MessageInvalidID = "The ID parameter is invalid."

	SQLXDriverName = "postgres"

	Canceled   types.OrderStatus = "canceled"
	Completed  types.OrderStatus = "completed"
	Paid       types.OrderStatus = "paid"
	Prepayment types.OrderStatus = "prepayment"
	Waiting    types.OrderStatus = "waiting"

	TableOrders               = "orders"
	ColumnIDTableOrders       = "id"
	RelationAgencyTableOrders = "Agency"
)
