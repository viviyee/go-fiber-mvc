package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/viviyee/go-mvc/controllers"
)

func PostRoutes(app *fiber.App) {
	app.Get("/posts", controllers.PostIndex)
}
