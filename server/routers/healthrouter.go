package routers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/shivam-jainn/constant/server/handlers"
)

func HealthRouter(app *fiber.App) {
	app.Get("/health", handlers.HealthyHandler)
}
