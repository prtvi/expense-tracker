package model

type Transaction struct {
	ID     string  `bson:"_id" json:"_id"`
	Date   string  `bson:"date" json:"date"`     // see this
	Desc   string  `bson:"desc" json:"desc"`     // see this
	Amount float32 `bson:"amount" json:"amount"` // see this
	Type   string  `bson:"type" json:"type"`     // see this with color
	Mode   string  `bson:"mode" json:"mode"`
	PaidTo string  `bson:"paid_to" json:"paid_to"`
}
