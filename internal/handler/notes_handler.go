package handler

import (
	"log"
	"net/http"

	"github.com/Pritam-25/go_crud_api_with_gin/internal/dto"
	"github.com/Pritam-25/go_crud_api_with_gin/internal/service"
	"github.com/gin-gonic/gin"
)

type NotesHandler struct {
	service *service.NoteService
}

func NewNotesHandler(service *service.NoteService) *NotesHandler {
	return &NotesHandler{
		service: service,
	}
}

func (h *NotesHandler) GetNotes(c *gin.Context) {
	notes, err := h.service.GetNotes(c.Request.Context())
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
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "invalid request body", // err.Error() can be used for more detailed error messages
		})
		return
	}

	note, err := h.service.CreateNote(c.Request.Context(), req)
	if err != nil {
		log.Printf("Error creating note: %v", err)
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
