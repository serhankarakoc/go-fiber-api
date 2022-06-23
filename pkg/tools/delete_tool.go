package tools

import (
	"github.com/gofiber/fiber/v2"
	"github.com/serhankarakoc/go-fiber-api/common/models"
)

func (h handler) DeleteTool(c *fiber.Ctx) error {
	id := c.Params("id")

	var tool models.Tool

	if result := h.DB.First(&tool, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	h.DB.Delete(&tool)

	return c.SendStatus(fiber.StatusOK)
}
