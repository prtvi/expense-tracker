package model

import "time"

type Transaction struct {
	ID     string    `bson:"_id" json:"_id"`
	Date   time.Time `bson:"date" json:"date"`
	Desc   string    `bson:"desc" json:"desc"`
	Amount float32   `bson:"amount" json:"amount"`
	Type   string    `bson:"type" json:"type"`
	Mode   string    `bson:"mode" json:"mode"`
	PaidTo string    `bson:"paid_to" json:"paid_to"`
}

// for returning a formatted transaction with date object converted to string
type TransactionFormatted struct {
	ID     string  `bson:"_id" json:"_id"`
	Date   string  `bson:"date" json:"date"`
	Desc   string  `bson:"desc" json:"desc"`
	Amount float32 `bson:"amount" json:"amount"`
	Type   string  `bson:"type" json:"type"`
	Mode   string  `bson:"mode" json:"mode"`
	PaidTo string  `bson:"paid_to" json:"paid_to"`
}
