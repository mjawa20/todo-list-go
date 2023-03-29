package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mjawa20/todo-list-go.git/db"
	_activityHandler "github.com/mjawa20/todo-list-go.git/modules/activity/handler"
	_activityRepository "github.com/mjawa20/todo-list-go.git/modules/activity/repository"
	_activityUsecase "github.com/mjawa20/todo-list-go.git/modules/activity/usecase"
	_todoHandler "github.com/mjawa20/todo-list-go.git/modules/todo/handler"
	_todoRepository "github.com/mjawa20/todo-list-go.git/modules/todo/repository"
	_todoUsecase "github.com/mjawa20/todo-list-go.git/modules/todo/usecase"
)

func Setup(f *fiber.App) {
	db := db.NewPostgres()

	activityRepository := _activityRepository.NewActivityRepository(db)
	activityUsecase := _activityUsecase.NewActivityUsecase(activityRepository)
	todoRepository := _todoRepository.NewTodoRepository(db)
	todoUsecase := _todoUsecase.NewTodoUsecase(todoRepository)

	_activityHandler.NewActivityHandler(f, activityUsecase)
	_todoHandler.NewTodoHandler(f, todoUsecase)
}
