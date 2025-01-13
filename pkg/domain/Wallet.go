package domain

import "time"

type Wallet struct {
	Name      string     `json:"name" bson:"name"`
	Month     time.Month `json:"month" bson:"month"`
	Revenues  []Revenues `json:"revenues" bson:"revenues"`
	Costs     []Costs    `json:"costs" bson:"costs"`
	CreatedAt time.Time  `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" bson:"updated_at"`
}
