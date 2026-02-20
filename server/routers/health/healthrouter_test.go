package health_test

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v3"
	health "github.com/shivam-jainn/constant/server/routers/health"
)

func TestHealthEndpoint(t *testing.T) {
	app := fiber.New()
	api := app.Group("/api")
	v1 := api.Group("/v1")
	health.HealthRouterHandler(api)

	req := httptest.NewRequest("GET", "/api/health/", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	if resp.StatusCode != 200 {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}
}
