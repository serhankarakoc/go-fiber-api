package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/serhankarakoc/go-fiber-api/common/config"
	"github.com/serhankarakoc/go-fiber-api/common/db"
	"github.com/serhankarakoc/go-fiber-api/pkg/tools"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed", err)
	}

	app := fiber.New()
	db := db.Init(c.DBUrl)

	tools.RegisterRoutes(app, db)

	app.Listen(c.Port)
}
