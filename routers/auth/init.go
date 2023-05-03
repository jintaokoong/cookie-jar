package auth

import (
	"github.com/gofiber/fiber/v2"
)

func Authenticate(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    "token_value",
		SameSite: "lax",
		HTTPOnly: true,
	})
	return c.JSON(fiber.Map{
		"status": "ok",
	})
}

func RemoveSession(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    "",
		SameSite: "lax",
		MaxAge:   -1,
		HTTPOnly: true,
	})
	return c.JSON(fiber.Map{
		"status": "ok",
	})
}
