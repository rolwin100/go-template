package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rolwin100/adopler/app/core/auth"
	"github.com/rolwin100/adopler/app/models"
	"github.com/rolwin100/adopler/app/utils"
)

func (controller *Controllers) Login(c *fiber.Ctx) error {

	return c.JSON(&fiber.Map{
		"success": "hello world",
	})
}

func (controller *Controllers) SignUp(c *fiber.Ctx) error {
	// create a new user struct
	signup := &models.SignUp{}
	if err := c.BodyParser(signup); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Create a new validator for a User model.
	validate := utils.NewValidator()

	if err := validate.Struct(signup); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Create a new user.
	auth := auth.Auth{}
	err := auth.SignUp(signup.Email, signup.Password, signup.Name)

	if err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	// Return status 200 and success message
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"error": false,
		"msg":   "User created successfully.",
	})
}
