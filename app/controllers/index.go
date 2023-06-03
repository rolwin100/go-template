package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rolwin100/adopler/app/core/index"
	"github.com/rolwin100/adopler/app/models"
)

func (controller Controllers) CreateIndex(c *fiber.Ctx) error {
	// create index core
	index := &index.Index{}
	indexDataModel := &models.Index{}

	// parse request body
	created, err := index.CreateIndex(indexDataModel.IndexName)

	if err != nil {
		return c.JSON(&fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	return c.JSON(&fiber.Map{
		"error":   false,
		"message": created,
	})
}
