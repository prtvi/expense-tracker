package model

type Transaction struct {
	Desc     string  `json:"desc"`
	To       string  `json:"to"`
	Date     string  `json:"date"`
	Amount   float32 `json:"amount"`
	Mode     string  `json:"mode"`
	PaidBy   string  `json:"paidBy"`
	Incoming bool    `json:"incoming"`
}

type CurrentBalance struct {
	Date   string  `json:"date"`
	Amount float32 `json:"amount"`
}
