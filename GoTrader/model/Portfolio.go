package model

type AccountType uint8

const (
	fund AccountType = iota
	stock
	index
	crypto
)

type BrokerType uint8

const (
	fund BrokerType = iota
	stock
	index
	crypto
)

type Portfolio struct {
	PortfolioId string
	accounts    map[AccountType]Account
	brokers     map[BrokerType]Broker
}
