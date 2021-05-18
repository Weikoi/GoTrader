package model

import "fmt"

type FundAccount struct {
	Account
}

func (fundAccount *FundAccount) Update() {
	fmt.Println("fund account update...")
}
