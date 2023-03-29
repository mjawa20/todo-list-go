package usecase

import "github.com/mjawa20/todo-list-go.git/domain"

type todoUsecase struct {
	repository domain.TodoRepository
}

func NewTodoUsecase(repository domain.TodoRepository) domain.TodoUseCase {
	return &todoUsecase{
		repository: repository,
	}
}

func (u *todoUsecase) GetAll(id uint) []domain.Todo {
	return u.repository.GetAll(id)
}

func (u *todoUsecase) GetByID(id uint) domain.Todo {
	return u.repository.GetByID(id)
}

func (u *todoUsecase) Create(todo *domain.Todo) error {
	return u.repository.Create(todo)
}

func (u *todoUsecase) Update(todo *domain.Todo) error {
	return u.repository.Update(todo)
}

func (u *todoUsecase) Delete(id uint) error {
	return u.repository.Delete(id)
}
