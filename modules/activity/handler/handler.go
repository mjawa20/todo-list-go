package handler

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/mjawa20/todo-list-go.git/domain"
)

type activityHandler struct {
	activityUseCase domain.ActivityUseCase
}

func NewActivityHandler(r fiber.Router, activityUseCase domain.ActivityUseCase) {
	handler := &activityHandler{
		activityUseCase: activityUseCase,
	}

	r.Get("/activity-groups", handler.GetAll)
	r.Get("/activity-groups/:id", handler.GetByID)
	r.Post("/activity-groups", handler.Create)
	r.Patch("/activity-groups/:id", handler.Update)
	r.Delete("/activity-groups/:id", handler.Delete)
}

func (h *activityHandler) GetAll(c *fiber.Ctx) error {
	activities := h.activityUseCase.GetAll()
	return domain.ResponseBuilder(c, "Success", 200, "Success", activities)
}

func (h *activityHandler) GetByID(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		fmt.Println(err)
	}

	activity := h.activityUseCase.GetByID(uint(id))
	if activity.Id == 0 {
		return domain.ResponseBuilder(c, "Not Found", 404, "Activity with Id "+strconv.Itoa(int(id))+" Not Found", nil)
	}

	return domain.ResponseBuilder(c, "Success", 200, "Success", activity)
}

func (h *activityHandler) Create(c *fiber.Ctx) error {
	activity := new(domain.Activities)
	if err := c.BodyParser(activity); err != nil {
		return err
	}

	if (*activity == domain.Activities{}) {
		return domain.ResponseBuilder(c, "Bad Request", 400, "title cannot be null", nil)
	}

	err := h.activityUseCase.Create(activity)
	if err != nil {
		return domain.ResponseBuilder(c, "Error", 404, err.Error(), nil)
	}
	return domain.ResponseBuilder(c, "Success", 200, "Success", activity)
}

func (h *activityHandler) Update(c *fiber.Ctx) error {
	id, errID := strconv.ParseUint(c.Params("id"), 10, 32)
	if errID != nil {
		fmt.Println(errID)
	}

	activity := new(domain.Activities)
	if err := c.BodyParser(activity); err != nil {
		return err
	}

	if (*activity == domain.Activities{}) {
		return domain.ResponseBuilder(c, "Bad Request", 400, "title cannot be null", nil)
	}

	newActivity, err := h.activityUseCase.Update(uint(id), activity)
	if err != nil {
		if err.Error() == "data not found" {
			return domain.ResponseBuilder(c, "Not Found", 404, "Activity with Id "+strconv.Itoa(int(id))+" Not Found", nil)
		}
		return domain.ResponseBuilder(c, "Error", 500, err.Error(), nil)
	}
	return domain.ResponseBuilder(c, "Success", 200, "Success", newActivity)

}

func (h *activityHandler) Delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		fmt.Println(err)
	}

	err = h.activityUseCase.Delete(uint(id))
	if err != nil {
		if err.Error() == "data not found" {
			return domain.ResponseBuilder(c, "Not Found", 404, "Activity with Id "+strconv.Itoa(int(id))+" Not Found", nil)
		}
		return domain.ResponseBuilder(c, "Error", 500, err.Error(), nil)
	}

	return domain.ResponseBuilder(c, "Success", 200, "activity was deleted", domain.Activities{})
}
