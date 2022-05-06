package model

type Summary struct {
	TotalIncome    float32 `bson:"total_income" json:"total_income"`
	TotalExpense   float32 `bson:"total_expense" json:"total_expense"`
	CurrentBalance float32 `bson:"current_balance" json:"current_balance"`
}
