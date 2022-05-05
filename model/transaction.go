package model

type Transaction struct {
	Date   string  `bson:"date"`
	Desc   string  `bson:"desc"`
	Amount float32 `bson:"amount"`
	Mode   string  `bson:"mode"`
	PaidTo string  `bson:"paid_to"`
	Type   string  `bson:"type"`
}
