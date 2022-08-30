package model

type ModeValues struct {
	Value     string `bson:"value" json:"value"`
	IsChecked bool   `bson:"is_checked" json:"is_checked"`
}

type Settings struct {
	Currency           string                `bson:"currency" json:"currency"`
	CurrentMonthBudget float32               `bson:"current_month_budget" json:"current_month_budget"`
	ModesOfPayment     map[string]ModeValues `bson:"modes_of_payment" json:"modes_of_payment"`
}
