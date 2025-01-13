package domain

import "time"

type Accounts struct {
	Id        string    `json:"id,omitempty" bson:"_id,omitempty"`
	Account   string    `json:"account,omitempty" bson:"account,omitempty"`
	Color     string    `json:"color,omitempty" bson:"color,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
