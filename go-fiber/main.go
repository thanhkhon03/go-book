package main

import (
	"github.com/gofiber/fiber/v2" //Gọi thư viện fiber
)

func main() {
	//Khởi tạo ứng dụng fiber
	app := fiber.New()

	// Định nghĩa route cơ bản
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello,world")
	})
	//Khởi động server tại cổng 5432
	app.Listen(":5432")

}
