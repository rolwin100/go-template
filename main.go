package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/rolwin100/adopler/app/controllers"
	"github.com/rolwin100/adopler/app/routes"
)

func main() {
	app := fiber.New()

	controllers := &controllers.Controllers{}
	routes.Routes(app, controllers)

	log.Fatal(app.Listen(":8000"))
}
