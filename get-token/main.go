package main

import (
	"context"
	"get-token/jwtutil"
	"get-token/repository"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Dependency injection struct
type AppDependencies struct {
	UserRepo repository.UserRepository
	JWTUtil  jwtutil.JWTUtil
}

func main() {
	// Kết nối đến MongoDB
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	db := client.Database("get-token")

	// Tạo repository và utils
	userRepo := repository.NewMongoUserRepository(db.Collection("users"))
	jwtUtil := jwtutil.NewJWTUtil("my_secret_key")

	// Khởi tạo dependencies
	deps := AppDependencies{
		UserRepo: userRepo,
		JWTUtil:  jwtUtil,
	}

	// Tạo ứng dụng Fiber
	app := fiber.New()
	app.Use(logger.New())

	// Route đăng ký
	app.Post("/register", deps.RegisterHandler)

	// Route đăng nhập
	app.Post("/login", deps.LoginHandler)

	// Route bảo vệ
	app.Get("/protected", deps.JWTMiddleware, func(c *fiber.Ctx) error {
		return c.SendString("Welcome to the protected route!")
	})

	// Chạy server
	log.Fatal(app.Listen(":8080"))
}
