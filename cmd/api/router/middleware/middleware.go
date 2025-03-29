package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func GetProtected(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Missing or invalid token",
		})
	}
	fmt.Println(token)
	return c.JSON(fiber.Map{"message": "You are authorized!"})
}
