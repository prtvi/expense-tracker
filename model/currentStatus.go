package model

type CurrentStatus struct {
	TotalIncome    float32 `bson:"total_income"`
	TotalExpense   float32 `bson:"total_expense"`
	CurrentBalance float32 `bson:"current_balance"`
}
