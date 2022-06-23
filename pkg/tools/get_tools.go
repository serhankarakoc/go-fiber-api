package tools

import (
	"github.com/gofiber/fiber/v2"
	"github.com/serhankarakoc/go-fiber-api/common/models"
)

func (h handler) GetTools(c *fiber.Ctx) error {
	var tools []models.Tool

	if result := h.DB.Find(&tools); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&tools)
}
