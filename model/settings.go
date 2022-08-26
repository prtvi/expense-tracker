package model

type Settings struct {
	Currency           string   `bson:"currency" json:"currency"`
	ModesOfPayment     []string `bson:"modes_of_payment" json:"modes_of_payment"`
	CurrentMonthBudget float32  `bson:"current_month_budget" json:"current_month_budget"`
}
