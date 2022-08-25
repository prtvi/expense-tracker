package model

type Summary struct {
	TotalIncome  float32 `bson:"total_income" json:"total_income"`
	TotalExpense float32 `bson:"total_expense" json:"total_expense"`
	TotalBalance float32 `bson:"total_balance" json:"total_balance"`
}
