package structs

import "time"

type Costs struct {
	Cost         string     `json:"cost" bson:"cost"`
	Amount       Money      `json:"amount" bson:"amount"`
	Installments *int       `json:"installments" bson:"installments"`
	Payment      Methods    `json:"payment" bson:"payment"`
	Category     Categories `json:"category" bson:"category"`
	CreatedAt    time.Time  `json:"created_at" bson:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at" bson:"updated_at"`
}
