package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/Pritam-25/go_crud_api_with_gin/internal/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type NoteRepository struct {
	collection *mongo.Collection
}

func NewNoteRepository(db *mongo.Database) *NoteRepository {
	return &NoteRepository{
		collection: db.Collection("notes"),
	}
}

func (r *NoteRepository) GetAll(ctx context.Context) ([]models.Note, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("repository: get all notes: %w", err)
	}

	var notes []models.Note

	if err := cursor.All(ctx, &notes); err != nil {
		return nil, fmt.Errorf("repository: decode notes: %w", err)
	}

	return notes, nil
}

func (r *NoteRepository) Create(ctx context.Context, note *models.Note) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	now := time.Now().UTC()
	note.CreatedAt = now
	note.UpdatedAt = now

	result, err := r.collection.InsertOne(ctx, note)
	if err != nil {
		return fmt.Errorf("repository: create note: %w", err)
	}

	if id, ok := result.InsertedID.(bson.ObjectID); ok {
		note.ID = id
		return nil
	}

	return fmt.Errorf("repository: failed to parse inserted id")
}
