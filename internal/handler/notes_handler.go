package handler

import (
	"net/http"

	"github.com/Pritam-25/go_crud_api_with_gin/internal/dto"
	"github.com/Pritam-25/go_crud_api_with_gin/internal/service"
	"github.com/gin-gonic/gin"
)

type NotesHandler struct {
	svc *service.NoteService
}

func NewNotesHandler(svc *service.NoteService) *NotesHandler {
	return &NotesHandler{
		svc: svc,
	}
}

func (h *NotesHandler) GetNotes(c *gin.Context) {
	notes, err := h.svc.GetNotes(c.Request.Context())
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
	var req dto.CreateNoteRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "invalid request body",
		})
		return
	}

	note, err := h.svc.CreateNote(c.Request.Context(), req)
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
