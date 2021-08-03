package formatter

type Formatter struct {
	OrderID        string `json:"order_id"`
	Fullname       string `json:"fullname"`
	IDCard         string `json:"id_card"`
	Car            string `json:"car"`
	CarType        string `json:"car_type"`
	EstimatedPrice int    `json:"estimated_price"`
}

type DoneFormatter struct {
	OrderID    string `json:"order_id"`
	Fullname   string `json:"fullname"`
	IDCard     string `json:"id_card"`
	Car        string `json:"car"`
	CarType    string `json:"car_type"`
	TotalPrice int    `json:"total_price"`
}
