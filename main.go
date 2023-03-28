package main

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/mjawa20/todo-list-go.git/routes"
)

func init() {
	os.Setenv("TZ", "Asia/Jakarta")
	loc, _ := time.LoadLocation("Asia/Jakarta")
	// handle err
	time.Local = loc // -> this is setting the global timezone

	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
}

func main() {
	f := fiber.New()
	routes.Setup(f)
	f.Listen(":8080")
}
