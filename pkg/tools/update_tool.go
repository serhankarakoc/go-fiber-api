package tools

import (
	"github.com/gofiber/fiber/v2"
	"github.com/serhankarakoc/go-fiber-api/common/models"
)

type UpdateToolRequestBody struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (h handler) UpdateTool(c *fiber.Ctx) error {
	id := c.Params("id")
	body := UpdateToolRequestBody{}

	// getting request's body
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var tool models.Tool

	if result := h.DB.First(&tool, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	tool.Title = body.Title
	tool.Description = body.Description

	// save tool
	h.DB.Save(&tool)

	return c.Status(fiber.StatusOK).JSON(&tool)
}
