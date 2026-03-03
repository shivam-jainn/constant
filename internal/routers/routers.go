package routers

import (
	"github.com/gofiber/fiber/v3"
	health "github.com/shivam-jainn/constant/internal/routers/health"
	orders "github.com/shivam-jainn/constant/internal/routers/orders"
)

func RegisterRouters(app *fiber.App) {
	apiGroup := app.Group("/api")

	v1ApiGroup := apiGroup.Group("/v1")

	health.HealthRouterHandler(apiGroup)
	orders.OrderRouterHandler(v1ApiGroup)
}
