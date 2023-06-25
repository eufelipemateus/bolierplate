package controllers

import "github.com/gofiber/fiber/v2"

func ExamplePage(c *fiber.Ctx) error {

	return c.JSON(&fiber.Map{
		"status":  "success",
		"message": "This is working",
	})
}
