package service

import (
	"context"

	"github.com/Pritam-25/go_crud_api_with_gin/internal/models"
	"github.com/Pritam-25/go_crud_api_with_gin/internal/repository"
)

// NoteService has a field called repo
type NoteService struct {
	repo *repository.NoteRepository
}

// NewNoteService is a constructor function that takes a NoteRepository and returns a NoteService
func NewNoteService(repo *repository.NoteRepository) *NoteService {
	return &NoteService{repo: repo}
}

func (s *NoteService) GetNotes(ctx context.Context) ([]models.Note, error) {
	return s.repo.GetAllNotes(ctx)
}

func (s *NoteService) CreateNote(ctx context.Context, req models.CreateNoteRequest) (*models.Note, error) {
	return s.repo.CreateNote(ctx, req)
}
