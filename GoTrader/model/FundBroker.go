package model

import "fmt"

type FundBroker struct {
	Broker
}

func (fundBroker *FundBroker) SubmitOrder() {
	fmt.Println("broker submit order")
}
