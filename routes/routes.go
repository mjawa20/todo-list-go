package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mjawa20/todo-list-go.git/db"
	_appHandler "github.com/mjawa20/todo-list-go.git/modules/activity/handler"
	_appRepository "github.com/mjawa20/todo-list-go.git/modules/activity/repository"
	_appUsecase "github.com/mjawa20/todo-list-go.git/modules/activity/usecase"
)

func Setup(f *fiber.App) {
	db := db.NewPostgres()

	appRepository := _appRepository.NewActivityRepository(db)
	appUsecase := _appUsecase.NewActivityUsecase(appRepository)

	_appHandler.NewActivityHandler(f, appUsecase)
}
