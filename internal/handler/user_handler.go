package handler

import (
	"github.com/Pritam-25/go_crud_api_with_gin/internal/service"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: service}
}
