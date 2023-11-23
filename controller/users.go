package controller

import (
	"github.com/gofiber/fiber/v2"
)

type User struct {
	Fisrtname string
	Lasname   string
	Age int32
}

func HandleUser(c *fiber.Ctx) error {
	user := User{
		Fisrtname: "Jake",
		Lasname:   "Sullyvan",
		Age:23,
	}
	return c.Status(fiber.StatusOK).JSON(user)
}

func HandleReciveUser(c *fiber.Ctx) error {
	user := User{}
	if err := c.BodyParser(&user); err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).JSON(user)
}
