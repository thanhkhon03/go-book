package repository

import (
	"github.com/gofiber/fiber/v2"
)

func (d AppDependencies) RegisterHandler(c *fiber.Ctx) error {
	var user User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
	}

	if err := d.UserRepo.CreateUser(c.Context(), &user); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to create user")
	}
	return c.SendString("User registered successfully")
}

func (d AppDependencies) LoginHandler(c *fiber.Ctx) error {
	var loginData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.BodyParser(&loginData); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
	}

	user, err := d.UserRepo.FindByUsername(c.Context(), loginData.Username)
	if err != nil || user.Password != loginData.Password {
		return c.Status(fiber.StatusUnauthorized).SendString("Invalid username or password")
	}

	token, err := d.JWTUtil.CreateToken(user.Username)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to create token")
	}

	return c.JSON(fiber.Map{
		"token": token,
	})
}

func (d AppDependencies) JWTMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).SendString("Missing token")
	}

	tokenString := authHeader[len("Bearer "):]
	_, err := d.JWTUtil.VerifyToken(tokenString)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).SendString("Invalid token")
	}

	return c.Next()
}
