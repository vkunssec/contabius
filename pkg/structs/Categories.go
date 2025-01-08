package structs

import "time"

type Categories struct {
	Grade     string    `json:"grade" bson:"grade"`
	Parent    *int      `json:"parent" bson:"parent"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}
