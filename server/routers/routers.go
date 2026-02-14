package routers

import "github.com/gofiber/fiber/v3"

func RegisterRouters(app *fiber.App) {
	HealthRouter(app)
}
