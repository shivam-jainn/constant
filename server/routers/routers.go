package routers

import (
	"github.com/gofiber/fiber/v3"
	health "github.com/shivam-jainn/constant/server/routers/health"
	orders "github.com/shivam-jainn/constant/server/routers/orders"
)

func RegisterRouters(app *fiber.App) {
	apiGroup := app.Group("/api")

	v1ApiGroup := apiGroup.Group("/v1")

	health.HealthRouterHandler(apiGroup)
	orders.OrderRouterHandler(v1ApiGroup)
}
