package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rolwin100/adopler/app/controllers"
)

func Routes(app *fiber.App, controllers *controllers.Controllers) {
	api := app.Group("/api")

	helloWorld := api.Group("/hello-world")
	helloWorld.Get("/", controllers.HelloWorld)

	/** Auth Routes */
	authRoute := api.Group("/auth")
	authRoute.Post("/login", controllers.Login)
	authRoute.Post("/signup", controllers.SignUp)

	/** Index Routes */
	indexRoute := api.Group("/index")
	indexRoute.Post("/", controllers.CreateIndex)
}
