package controllers

import "github.com/gofiber/fiber/v2"

func (controller Controllers) HelloWorld(c *fiber.Ctx) error {

	return c.JSON(&fiber.Map{
		"success": "hello world",
	})

}
