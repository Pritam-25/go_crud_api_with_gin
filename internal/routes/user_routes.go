package routes

import (
	"github.com/Pritam-25/go_crud_api_with_gin/internal/handler"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(rg gin.IRouter, userHandler *handler.UserHandler) {
	// user := rg.Group("/user")
	// {
	// 	// user.POST("/{id}", userHandler.Register)
	// 	// user.POST("/login", userHandler.Login) // Implement Login handler in UserHandler
	// }
}
