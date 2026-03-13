package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/Pritam-25/go_crud_api_with_gin/internal/dto"
	"github.com/Pritam-25/go_crud_api_with_gin/internal/models"
	"github.com/Pritam-25/go_crud_api_with_gin/internal/repository"
	"github.com/Pritam-25/go_crud_api_with_gin/internal/utils"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(ctx context.Context, req dto.RegisterRequest) (*models.User, error) {
	// check if user with the same email already exists
	_, err := s.repo.FindByEmail(ctx, req.Email)
	if err == nil {
		return nil, fmt.Errorf("user with this email already exists")
	}

	if !errors.Is(err, mongo.ErrNoDocuments) {
		return nil, err
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password")
	}

	user := &models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
		Role:     req.Role,
	}

	err = s.repo.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
