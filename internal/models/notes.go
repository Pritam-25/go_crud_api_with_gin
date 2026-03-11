package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Note struct {
	ID        	bson.ObjectID      `json:"id" bson:"_id,omitempty"`
	Title     	string 			   `json:"title" bson:"title"`
	Content   	string             `json:"content" bson:"content"`
	Pinned    	bool               `json:"pinned" bson:"pinned"`
	CreatedAt 	time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt 	time.Time          `json:"updated_at" bson:"updated_at"`
}

type CreateNoteRequest struct {
	Title    string `json:"title" binding:"required"`
	Content  string `json:"content" binding:"required"`
	Pinned   bool   `json:"pinned"`
}