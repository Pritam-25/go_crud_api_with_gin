package server

import (
	"net/http"

	"github.com/Pritam-25/go_crud_api_with_gin/internal/handler"
	"github.com/Pritam-25/go_crud_api_with_gin/internal/repository"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func NewRouter(db *mongo.Database) *gin.Engine {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Welcome to Go CRUD API",
		})
	})

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	notRepo := repository.NewNoteRepository(db)
	noteHandler := handler.NewNotesHandler(notRepo)

	router.GET("/notes", noteHandler.GetNotes)
	router.POST("/notes", noteHandler.CreateNote)

	return router
}