package main

import (
	"cookiejar/config"
	"cookiejar/routers/auth"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// Initialize configuration
	config.Init()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,                    // we need this to allow the cookie to be sent
		AllowOrigins:     "http://localhost:5173", // whitelist the frontend domain
	}))

	// declare health check route
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
		})
	})

	// we set the cookie on signin
	app.Post("/api/auth/signin", auth.Authenticate)

	// we destroy the cookie on signout
	app.Post("/api/auth/signout", auth.RemoveSession)

	// we define a route that is required to have a valid refresh token
	app.Get("/api/secure/resource", func(c *fiber.Ctx) error {
		cookie := c.Cookies("refresh_token")
		if cookie == "token_value" {
			return c.Next()
		}
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status": "unauthorized",
		})
	}, func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
		})
	})

	app.Listen(":3000")
}
