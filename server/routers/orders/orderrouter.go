package orders

import (
	"github.com/gofiber/fiber/v3"
	"github.com/shivam-jainn/constant/server/handlers"
)

func OrderRouterHandler(apiGroup fiber.Router) {
	OrderRouter := apiGroup.Group("/order")
	OrderRouter.Post("/place", handlers.OrderPlaceHandler)
}
