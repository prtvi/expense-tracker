package model

type Budget struct {
	Budget    float32 `bson:"budget" json:"budget"`
	Month     int     `bson:"month" json:"month"`
	Year      int     `bson:"year" json:"year"`
	Spent     float32 `bson:"spent" json:"spent"`
	Remaining float32 `bson:"remaining" json:"remaining"`
}
