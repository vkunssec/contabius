package domain

import "time"

type Revenues struct {
	Revenue   string    `json:"revenue" bson:"revenue"`
	Amount    Money     `json:"amount" bson:"amount"`
	Method    *Methods  `json:"method" bson:"method"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}
