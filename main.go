package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/mjawa20/todo-list-go.git/routes"
)

func init() {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	time.Local = loc // -> this is setting the global timezone

	godotenv.Load()
}

func main() {
	f := fiber.New()
	routes.Setup(f)
	f.Listen(":8090")
}
