package domain

import "time"

type Accounts struct {
	Id        string    `json:"id,omitempty" bson:"_id,omitempty" example:"678079f6f5080a39a8eedc1e"`
	Account   string    `json:"account,omitempty" bson:"account,omitempty" example:"Conta Corrente"`
	Color     string    `json:"color,omitempty" bson:"color,omitempty" example:"#000000"`
	CreatedAt time.Time `json:"created_at,omitempty" bson:"created_at,omitempty" example:"2025-01-01T00:00:00Z"`
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty" example:"2025-01-01T00:00:00Z"`
}
