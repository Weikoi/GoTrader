package model

import "fmt"

type FundAccount struct {
	AccountId string
	Positions
}

func (fundAccount *FundAccount) Update() {
	fmt.Println("fund account update...")
}

func (fundAccount *FundAccount) ProcessDividend() {
	fmt.Println("fund account update...")
}

func (fundAccount *FundAccount) ProcessUnfilledOrders() {
	fmt.Println("fund account update...")
}
