package model

type Summary struct {
	Income  float32 `bson:"income" json:"income"`
	Expense float32 `bson:"expense" json:"expense"`
	Balance float32 `bson:"balance" json:"balance"`
}
