package main

import (
	"log"

	TC "GOTest/Controllers/testController"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	app.Get("/:imgName", TC.HandleGet)
	app.Post("/item", TC.HandlePost)
	log.Fatal(app.Listen(":3000"))
}
