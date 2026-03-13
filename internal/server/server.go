package server

import (
	"time"

	"github.com/Pritam-25/go_crud_api_with_gin/internal/handler"
	"github.com/Pritam-25/go_crud_api_with_gin/internal/middleware"
	"github.com/Pritam-25/go_crud_api_with_gin/internal/routes"
	"github.com/gin-gonic/gin"
)

func NewRouter(noteHandler *handler.NotesHandler, authHandler *handler.AuthHandler, userHandler *handler.UserHandler) *gin.Engine {
	router := gin.Default()
	router.Use(middleware.CORSMiddleware())
	router.Use(middleware.TimeoutMiddleware(5 * time.Second))

	api := router.Group("/api/v1")
	routes.RegisterHealthRoutes(router)
	routes.RegisterNoteRoutes(api, noteHandler)
	routes.RegisterAuthRoutes(api, authHandler)
	routes.RegisterUserRoutes(api, userHandler)

	return router
}
