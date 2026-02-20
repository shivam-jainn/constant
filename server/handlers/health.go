package handlers

import "github.com/gofiber/fiber/v3"

func HealthyHandler(c fiber.Ctx) error {
	return c.SendString("Healthy")
}
