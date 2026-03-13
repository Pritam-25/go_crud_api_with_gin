package routes

import (
	"github.com/Pritam-25/go_crud_api_with_gin/internal/handler"
	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(rg gin.IRouter, authHandler *handler.AuthHandler) {
	auth := rg.Group("/auth")
	{
		auth.POST("/register", authHandler.Register)
		// auth.POST("/login", authHandler.Login) // Implement Login handler in AuthHandler
	}
}
