package health

import (
	"github.com/gofiber/fiber/v3"
	"github.com/shivam-jainn/constant/server/handlers"
)

func HealthRouterHandler(apiGroup fiber.Router) {
	healthRouter := apiGroup.Group("/health")
	healthRouter.Get("/", handlers.HealthyHandler)
}
