package user

import (
	"github.com/Rolas444/apigo_base/cmd/api/dtos"
	"github.com/gofiber/fiber/v2"
)

func (h *handler) Create(c *fiber.Ctx) error {

	var userDTO dtos.UserCreateDTO
	if err := c.BodyParser(&userDTO); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Los datos no son válidos"})
	}
	user := userDTO.ToDomain()
	createdUser, err := h.Service.Create(c.Context(), user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	var userResponse dtos.UserDTO
	userResponse.FromDomain(*createdUser)

	return c.Status(fiber.StatusCreated).JSON(userResponse)

}
