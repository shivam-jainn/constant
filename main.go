package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/shivam-jainn/constant/server/routers"
)

func main() {

	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	routers.RegisterRouters(app)

	app.Listen(":3000")

}
