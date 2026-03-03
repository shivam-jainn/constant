package handlers

import "github.com/gofiber/fiber/v3"

type OrderRequest struct {
	Quantity int64 `json:"quantity"`
}

func OrderPlaceHandler(c fiber.Ctx) error {
	var orderRequest OrderRequest

	// Parse the JSON request body into the struct
	if err := c.Bind().JSON(&orderRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Use orderRequest.Quantity here
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":  "Order placed",
		"quantity": orderRequest.Quantity,
	})
}
