package app

import (
	"github.com/Pritam-25/go_crud_api_with_gin/internal/config"
	"github.com/Pritam-25/go_crud_api_with_gin/internal/db"
	"github.com/Pritam-25/go_crud_api_with_gin/internal/handler"
	"github.com/Pritam-25/go_crud_api_with_gin/internal/repository"
	"github.com/Pritam-25/go_crud_api_with_gin/internal/server"
	"github.com/Pritam-25/go_crud_api_with_gin/internal/service"
	"github.com/gin-gonic/gin"
)

func BuildServer(cfg *config.Config) (*gin.Engine, func(), error) {
	client, database, err := db.ConnectMongoDB(cfg)
	if err != nil {
		return nil, nil, err
	}

	// dependencies
	noteRepo := repository.NewNoteRepository(database)
	noteService := service.NewNoteService(noteRepo)
	noteHandler := handler.NewNotesHandler(noteService)

	router := server.NewRouter(noteHandler)

	cleanup := func() {
		db.DisconnectMongoDB(client)
	}

	return router, cleanup, nil
}
