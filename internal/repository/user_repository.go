package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/Pritam-25/go_crud_api_with_gin/internal/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) *UserRepository {
	return &UserRepository{
		collection: db.Collection("users"),
	}
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (models.User, error) {
	var user models.User

	filter := bson.M{"email": email}

	err := r.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return models.User{}, fmt.Errorf("repository: find user by email: %w", err)
	}

	return user, nil
}

func (r *UserRepository) Create(ctx context.Context, user *models.User) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	now := time.Now().UTC()
	user.CreatedAt = now
	user.UpdatedAt = now

	result, err := r.collection.InsertOne(ctx, user)
	if err != nil {
		return fmt.Errorf("repository: insert user: %w", err)
	}

	if id, ok := result.InsertedID.(bson.ObjectID); ok {
		user.ID = id
		return nil
	}

	return fmt.Errorf("repository: failed to parse inserted id")
}
