package domain

import "time"

type Investments struct {
	Investment string    `json:"investment" bson:"investment"`
	Amount     Money     `json:"amount" bson:"amount"`
	Account    Accounts  `json:"account" bson:"account"`
	CreatedAt  time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" bson:"updated_at"`
}
