package main

import (
	"log"

	"github.com/Pritam-25/go_crud_api_with_gin/internal/app"
	"github.com/Pritam-25/go_crud_api_with_gin/internal/config"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("failed to load config:", err)
	}

	router, cleanup, err := app.BuildServer(cfg)
	if err != nil {
		log.Fatal("failed to build app:", err)
	}

	defer cleanup()

	log.Printf("Server running on port http://localhost:%s", cfg.Port)

	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatal("failed to start server:", err)
	}
}
