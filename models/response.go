package models

type ResponseForFront struct{
	Brand string `json:"brand"`
	Places string `json:"places"`
	Carnum string `json:"carnum"`
	Date string `json:"date"`
	OrderId string `json:"order_id"`
	NowTime string `json:"now_time"`
	Token  string `json:"token"`
}