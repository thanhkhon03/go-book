package middleware

import (
    "fmt"
    "net/http"
)


func authMiddleware(c *fiber.Ctx) error {
	// Lấy token từ header Authorization
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).SendString("Missing or invalid token")
	}

	// Token có dạng: "Bearer <token>"
	tokenString := authHeader[len("Bearer "):]

	// Xác minh token
	_, err := verifyToken(tokenString)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).SendString("Invalid token")
	}

	// Nếu token hợp lệ, tiếp tục xử lý
	return c.Next()
}
