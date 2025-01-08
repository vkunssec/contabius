package structs

import "time"

type Accounts struct {
	Id        string    `json:"id,omitempty" bson:"_id,omitempty"`
	Account   string    `json:"account" bson:"account,omitempty"`
	Color     string    `json:"color" bson:"color,omitempty"`
	CreatedAt time.Time `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at,omitempty"`
}
