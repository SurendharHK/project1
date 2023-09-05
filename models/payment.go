package models

type Payments struct {
	Id          string       `json:"_id" bson:"_id"`
	Customer_id string       `json:"customer_id" bson:"customer_id"`
	Status      string       `json:"status" bson:"status"`
	Gateway     string       `json:"gateway" bson:"gateway"`
	Type        string       `json:"type" bson:"type"`
	Amount      float64      `json:"amount" bson:"amount"`
	Card        Paymentscard `json:"card" bson:"card"`
	Token       string       `json:"token" bson:"token"`
}

type Paymentscard struct {
	CardNo          float64   `json:"cardno" bson:"cardno"`
	Brand           string  `json:"brand" bson:"brand"`
	PanLastFourNo   string  `json:"panlastfourno" bson:"panlastfourno"`
	ExpirationMonth int16   `json:"expirationmonth" bson:"expirationmonth"`
	ExpirationYear  int16   `json:"expirationyear" bson:"expirationyear"`
	Cvvverified     int16   `json:"cvvverified" bson:"cvvverified"`
	Balance         float64 `json:"balance" bson:"balance"`
}

