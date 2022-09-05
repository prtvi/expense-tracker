package model

type Summary struct {
	TotalIncome  float32      `bson:"total_income" json:"total_income"`
	TotalExpense float32      `bson:"total_expense" json:"total_expense"`
	TotalBalance float32      `bson:"total_balance" json:"total_balance"`
	IndiModeSums IndiModeSums `bson:"indi_mode_sums" json:"indi_mode_sums"`
}

type Mode struct {
	Income  float32 `bson:"income" json:"income"`
	Expense float32 `bson:"expense" json:"expense"`
}

type IndiModeSums map[string]Mode
