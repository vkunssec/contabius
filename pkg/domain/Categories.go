package domain

import "time"

type Categories struct {
	Id        string    `json:"id,omitempty" bson:"_id,omitempty" example:"678079f6f5080a39a8eedc1e"`
	Category  string    `json:"category" bson:"category" example:"Alimentação"`
	Parent    *string   `json:"parent" bson:"parent" example:"678079f6f5080a39a8eedc1e"`
	CreatedAt time.Time `json:"created_at,omitempty" bson:"created_at,omitempty" example:"2025-01-01T00:00:00Z"`
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty" example:"2025-01-01T00:00:00Z"`
}
