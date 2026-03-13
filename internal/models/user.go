package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type User struct {
	ID        bson.ObjectID `json:"id" bson:"_id,omitempty"`
	Username  string        `json:"username" bson:"username"`
	Email     string        `json:"email" bson:"email"`
	Password  string        `json:"password" bson:"password"`
	Role      string        `json:"role" bson:"role"`
	CreatedAt time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time     `json:"updated_at" bson:"updated_at"`
}
