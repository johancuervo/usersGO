package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/johancuervo/usersGO/internal/domain"
	"github.com/johancuervo/usersGO/internal/service"
)

type UserHandler struct {
	service *service.UserService
}

// constructor
func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	user := new(domain.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}
	if err := h.service.CreateUser(user); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(201).JSON(user)
}
func (h *UserHandler) GetAllUsers(c *fiber.Ctx) error {
	users, err := h.service.GetAllUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(users)
}
func (h *UserHandler) FindUserById(c *fiber.Ctx) error {
	userId := c.Params("userId")
	users, err := h.service.GetUser(userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(users)
}

func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	userId := c.Params("userId")
	var updatedUser *domain.User
	if err := c.BodyParser(&updatedUser); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
	}

	err := h.service.UpdateUser(updatedUser, userId)
	if err != nil {
		if err.Error() == "user not found" {
			return c.Status(fiber.StatusNotFound).SendString("User Not Found")
		}
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(updatedUser)
}

func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	userId := c.Params("userId")
	err := h.service.DeleteUser(userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return nil
}
