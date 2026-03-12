package main

import (
	"log"

	"github.com/Pritam-25/go_crud_api_with_gin/internal/config"
	"github.com/Pritam-25/go_crud_api_with_gin/internal/db"
	"github.com/Pritam-25/go_crud_api_with_gin/internal/handler"
	"github.com/Pritam-25/go_crud_api_with_gin/internal/repository"
	"github.com/Pritam-25/go_crud_api_with_gin/internal/server"
	"github.com/Pritam-25/go_crud_api_with_gin/internal/service"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("failed to load config:", err)
	}

	client, database, err := db.ConnectMongoDB(cfg)
	if err != nil {
		log.Fatal("failed to connect to MongoDB:", err)
	}

	defer db.DisconnectMongoDB(client)

	// Wire dependencies
	noteRepo := repository.NewNoteRepository(database)
	noteService := service.NewNoteService(noteRepo)
	noteHandler := handler.NewNotesHandler(noteService)

	router := server.NewRouter(noteHandler)

	log.Printf("Server running on port http://localhost:%s", cfg.Port)

	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatal("failed to start server:", err)
	}
}
