package model

type ResponseMsg struct {
	StatusCode int  `bson:"status_code" json:"status_code"`
	Success    bool `bson:"success" json:"success"`
}
