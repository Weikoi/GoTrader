package config

/*
枚举类型
*/

// OrderStatus ===========================================
// OrderStatus 订单状态
type OrderStatus int

const (
	UNSUBMITTED OrderStatus = 1
	SUBMITTED   OrderStatus = 2
	FILLED      OrderStatus = 3
	CANCELED    OrderStatus = 4
	FAILED      OrderStatus = 5
)

// OrderType ============================================
// OrderType 订单类型
type OrderType int

const (
	PURCHASE OrderType = 1
	REDEEM   OrderType = 2
)
