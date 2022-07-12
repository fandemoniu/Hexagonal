package handlers

import (
	"hexagonal/internals/core/ports"

	"github.com/gofiber/fiber/v2"
)

type UserHandlers struct {
	userService ports.UserService
}

func NewUserHandlers(userService ports.UserService) *UserHandlers {
	return &UserHandlers{
		userService: userService,
	}
}

// Login implements ports.UserHandlers
func (h *UserHandlers) Login(c *fiber.Ctx) error {
	var email string
	var password string
	// Extract the body and get the email and password
	err := h.userService.Login(email, password)
	if err != nil {
		return err
	}
	return nil
}

// Register implements ports.UserHandlers
func (h *UserHandlers) Register(c *fiber.Ctx) error {
	var email string
	var password string
	var confirmPassword string

	// Extract the body and get the email and password
	err := h.userService.Register(email, password, confirmPassword)
	if err != nil {
		return err
	}
	return nil
}

var _ ports.UserHandlers = (*UserHandlers)(nil)
