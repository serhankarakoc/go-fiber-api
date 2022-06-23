package tools

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(app *fiber.App, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := app.Group("/tools")
	routes.Post("/", h.AddTool)
	routes.Get("/", h.GetTools)
	routes.Get("/:id", h.GetTool)
	routes.Put("/:id", h.UpdateTool)
	routes.Delete("/:id", h.DeleteTool)
}
