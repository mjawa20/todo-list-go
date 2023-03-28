package domain

import (
	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Status  string      `json:"status,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func ResponseBuilder(c *fiber.Ctx, status string, statusCode int, message string, data interface{}) error {
	return c.Status(statusCode).JSON(Response{Status: status, Message: message, Data: data})
}
