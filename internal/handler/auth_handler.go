package handler

import (
	"log"
	"net/http"

	"github.com/Pritam-25/go_crud_api_with_gin/internal/dto"
	"github.com/Pritam-25/go_crud_api_with_gin/internal/service"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service *service.UserService
}

func NewAuthHandler(service *service.UserService) *AuthHandler {
	return &AuthHandler{service: service}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("Error binding JSON: %v", err)

		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "invalid request body",
		})
		return
	}

	user, err := h.service.CreateUser(c.Request.Context(), req)
	if err != nil {
		log.Printf("Error creating user: %v", err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "failed to create user",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    user,
	})
}
