package handler

import (
	"net/http"

	"github.com/Pritam-25/go_crud_api_with_gin/internal/models"
	"github.com/Pritam-25/go_crud_api_with_gin/internal/repository"
	"github.com/gin-gonic/gin"
)

type NotesHandler struct {
	repo *repository.NoteRepository
}

func NewNotesHandler(repo *repository.NoteRepository) *NotesHandler {
	return &NotesHandler{
		repo: repo,
	}
}

func (h *NotesHandler) GetNotes(c *gin.Context) {
	notes, err := h.repo.GetAllNotes(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "failed to fetch notes",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    notes,
	})
}

func (h *NotesHandler) CreateNote(c *gin.Context) {
	var req models.CreateNoteRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "invalid request body",
		})
		return
	}

	note, err := h.repo.CreateNote(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "failed to create note",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    note,
	})
}