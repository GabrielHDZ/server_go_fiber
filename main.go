package main

import (
	"github.com/gabrielhdz/server_go_fiber/controller"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func main() {
	app := fiber.New()

	//Middleware
	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	//aplicate middleware headerId only the group
	app.Use(requestid.New())
	userGroup := app.Group("/users")
	userGroup.Get("", controller.HandleUser)
	userGroup.Post("", controller.HandleReciveUser)
	app.Listen(":3000")
}
