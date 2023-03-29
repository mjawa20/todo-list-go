package handler

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/mjawa20/todo-list-go.git/domain"
)

type todoHandler struct {
	todoUseCase domain.TodoUseCase
}

func NewTodoHandler(r fiber.Router, todoUseCase domain.TodoUseCase) {
	handler := &todoHandler{
		todoUseCase: todoUseCase,
	}

	r.Get("/todo-items", handler.GetAll)
	r.Get("/todo-items/:id", handler.GetByID)
	r.Post("/todo-items", handler.Create)
	r.Patch("/todo-items/:id", handler.Update)
	r.Delete("/todo-items/:id", handler.Delete)
}

func (h *todoHandler) GetAll(c *fiber.Ctx) error {
	activityId := c.Query("activity_group_id")
	intVar, err := strconv.ParseUint(activityId, 10, 32)
	if err != nil && activityId != "" {
		return domain.ResponseBuilder(c, "Error", 400, err.Error(), nil)
	}

	todos := h.todoUseCase.GetAll(uint(intVar))
	return domain.ResponseBuilder(c, "Success", 200, "Success", todos)
}

func (h *todoHandler) GetByID(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		fmt.Println(err)
	}

	todo := h.todoUseCase.GetByID(uint(id))

	if todo.Id == 0 {
		return domain.ResponseBuilder(c, "Not Found", 404, "Todo with Id "+strconv.Itoa(int(id))+" Not Found", nil)
	}

	return domain.ResponseBuilder(c, "Success", 200, "Success", todo)
}

func (h *todoHandler) Create(c *fiber.Ctx) error {
	todo := new(domain.Todos)
	if err := c.BodyParser(todo); err != nil {
		return err
	}

	if (*todo == domain.Todos{}) {
		return domain.ResponseBuilder(c, "Bad Request", 400, "title cannot be null", nil)
	}

	err := h.todoUseCase.Create(todo)
	if err != nil {
		return domain.ResponseBuilder(c, "Error", 400, err.Error(), nil)
	}
	return domain.ResponseBuilder(c, "Success", 200, "Success", todo)
}

func (h *todoHandler) Update(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		fmt.Println(err)
	}

	todo := new(domain.Todos)
	if err := c.BodyParser(todo); err != nil {
		return err
	}
	newTodo, err := h.todoUseCase.Update(uint(id), todo)
	if err != nil {
		if err.Error() == "data not found" {
			return domain.ResponseBuilder(c, "Not Found", 404, "Todo with Id "+strconv.Itoa(int(id))+" Not Found", nil)
		}
		return domain.ResponseBuilder(c, "Error", 500, err.Error(), nil)
	}
	return domain.ResponseBuilder(c, "Success", 200, "Success", newTodo)

}

func (h *todoHandler) Delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return domain.ResponseBuilder(c, "Error", 400, err.Error(), nil)
	}

	err = h.todoUseCase.Delete(uint(id))
	if err != nil {
		if err.Error() == "data not found" {
			return domain.ResponseBuilder(c, "Not Found", 404, "Todo with Id "+strconv.Itoa(int(id))+" Not Found", nil)
		}
		return domain.ResponseBuilder(c, "Error", 500, err.Error(), nil)
	}

	return domain.ResponseBuilder(c, "Success", 200, "todo was deleted", nil)
}
