package middleware

import (
	"errors"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func JWTAuthen() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Secret key for JWT
		secret := []byte(os.Getenv("JWT_SECRET_KEY"))

		// Get the Authorization header from the request
		header := c.Get("Authorization")
		if header == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"what":       "error",
				"error_code": "NoToken",
				"msg":        "No JWT token provided",
			})
		}

		// Remove "Bearer " from the token string
		tokenString := strings.Replace(header, "Bearer ", "", 1)

		// Parse and validate the token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Check if the signing method is HMAC
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				errMsg := errors.New("Unexpected signing method: " + token.Header["alg"].(string))
				return nil, errMsg
			}

			return secret, nil
		})

		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"what":       "error",
				"error_code": "InvalidToken",
				"msg":        err.Error(),
			})
		}

		// Check if the token is valid
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// Set the "userId" value to the fiber context
			c.Locals("userId", claims["userId"])
			return c.Next()
		}

		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"what":       "error",
			"error_code": "InvalidToken",
			"msg":        "JWT token is not valid",
		})
	}
}
