package server

import (
	"time"

	"github.com/Pritam-25/go_crud_api_with_gin/internal/handler"
	"github.com/Pritam-25/go_crud_api_with_gin/internal/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter(noteHandler *handler.NotesHandler) *gin.Engine {
	router := gin.New()

	router.Use(gin.Recovery())
	router.Use(middleware.RequestLogger())
	router.Use(middleware.CORSMiddleware())
	router.Use(middleware.TimeoutMiddleware(5 * time.Second))

	RegisterHealthRoutes(router)

	api := router.Group("/api/v1")
	RegisterNoteRoutes(api, noteHandler)

	return router
}
