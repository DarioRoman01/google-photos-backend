package middlewares

import (
	"strings"

	"github.com/DarioRoman01/photos/models"
	"github.com/DarioRoman01/photos/utils"
	"github.com/gofiber/fiber/v2"
)

var (
	NO_AUTH_NEEDED = []string{
		"login",
		"signup",
		"verify",
	}
)

func shoulCheckToken(route string) bool {
	for _, p := range NO_AUTH_NEEDED {
		if strings.Contains(route, p) {
			return false
		}
	}
	return true
}

func CheckAuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if !shoulCheckToken(c.Path()) {
			return c.Next()
		}

		token := c.Cookies("token")
		if token == "" {
			return c.Status(401).JSON(utils.JsonError("Unauthorized"))
		}

		claims, err := utils.VerifyToken(token)
		if err != nil {
			return c.Status(401).JSON(utils.JsonError("Unauthorized"))
		}

		if claims.Type != models.AccessToken.String() {
			return c.Status(401).JSON(utils.JsonError("Invalid token"))
		}

		c.Locals("user_id", claims.UserID)
		c.Locals("username", claims.Username)
		return c.Next()
	}
}
