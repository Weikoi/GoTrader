package model

import "time"
import "GoTrader/GoTrader/config"

type Order struct {
	// 订单编号
	OrderId int
	//订单状态
	Status config.OrderStatus
	//订单类型
	Type config.OrderType
	//订单提交时间
	SubmitTime time.Time
	//订单变动份额
	OccurShares int
	//订单变动资金
	OccurBalance int
}
