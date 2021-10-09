package model

import "time"

type Laptop struct {
	tableName struct{} `pg:"test.laptop"`
	Id        string   `pg:"id,pk"`
	Name      string
	Cpu       string
}

type Bike struct {
	tableName struct{} `pg:"test.bike"`
	Id        string   `pg:"id,pk"`
	Name      string
	Volume    int
}

type PriceHistory struct {
	tableName   struct{} `pg:"test.price_history"`
	Id          string   `pg:"id,pk"`
	ItemId      string
	Type        string
	Price       int
	CreatedDate time.Time `pg:"default:now()"`
}
