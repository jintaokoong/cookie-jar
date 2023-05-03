package auth

import (
	"github.com/gofiber/fiber/v2"
)

func Authenticate(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    "token_value",
		Domain:   "localhost", // setting domain to hoge.com will make the cookie accessible for both fuga.hoge.com and piyo.hoge.com
		SameSite: "strict",    // the domain setting allows us to set the SameSite attribute to strict, which is the most secure option
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
