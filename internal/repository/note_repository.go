package repository

import (
	"context"
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

func (r *NoteRepository) GetAllNotes(ctx context.Context) ([]models.Note, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var notes []models.Note

	if err := cursor.All(ctx, &notes); err != nil {
		return nil, err
	}

	return notes, nil
}

func (r *NoteRepository) CreateNote(ctx context.Context, req models.CreateNoteRequest) (*models.Note, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	now := time.Now().UTC()
	note := &models.Note{
		Title:     req.Title,
		Content:   req.Content,
		Pinned:    req.Pinned,
		CreatedAt: now,
		UpdatedAt: now,
	}

	result, err := r.collection.InsertOne(ctx, note)
	if err != nil {
		return nil, err
	}

	if insertedID, ok := result.InsertedID.(bson.ObjectID); ok {
		note.ID = insertedID
	}

	return note, nil
}
