package middleware

import (
	"github.com/RianIhsan/go-code-nexus/domain/users"
	"github.com/RianIhsan/go-code-nexus/utils/jwt"
	"github.com/RianIhsan/go-code-nexus/utils/response"
	"github.com/gofiber/fiber/v2"
	jwtMapClaims "github.com/golang-jwt/jwt"
	"log"
	"strings"
)

func Protected(jwtService jwt.IJwt, userService users.IServiceUser) fiber.Handler {
	return func(c *fiber.Ctx) error {
		header := c.Get("Authorization")

		if !strings.HasPrefix(header, "Bearer ") {
			return response.SendStatusUnauthorized(c, "Access denied: missing token")
		}

		tokenString := strings.TrimPrefix(header, "Bearer ")

		token, err := jwtService.ValidateToken(tokenString)
		if err != nil {
			return response.SendStatusUnauthorized(c, "Access denied: invalid token")
		}

		claim, ok := token.Claims.(jwtMapClaims.MapClaims)
		if !ok || !token.Valid {
			return response.SendStatusUnauthorized(c, "Access denied: invalid token")
		}

		userID := int(claim["user_id"].(float64))

		user, err := userService.GetId(userID)
		if err != nil {
			log.Printf("Error retrieving user: %v", err)
			return response.SendStatusUnauthorized(c, "Access denied: user not found")
		}

		c.Locals("CurrentUser", user)

		return c.Next()
	}
}
