package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/viviyee/go-mvc/appsetting"
	"github.com/viviyee/go-mvc/routes"
)

func init() {
	appsetting.LoadENV()
	appsetting.ConnectDB()
	appsetting.SyncDB()
}

func main() {
	/* app */
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Static("/", "./public")

	/* routes */
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	routes.PostRoutes(app)

	/* app run */
	app.Listen(":" + os.Getenv("PORT"))
}
