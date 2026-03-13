package routes

import (
	"github.com/Pritam-25/go_crud_api_with_gin/internal/handler"
	"github.com/gin-gonic/gin"
)

func RegisterNoteRoutes(rg gin.IRouter, noteHandler *handler.NotesHandler) {
	notes := rg.Group("/notes")
	{
		notes.GET("", noteHandler.GetNotes)
		notes.POST("", noteHandler.CreateNote)
	}
}
