package structs

import "time"

type Methods struct {
	Method    string    `json:"method" bson:"method"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}
