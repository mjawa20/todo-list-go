package repository

import (
	"errors"

	"github.com/mjawa20/todo-list-go.git/domain"
)

type todoRepository struct {
	db domain.DB
}

func NewTodoRepository(db domain.DB) domain.TodoRepository {
	return &todoRepository{
		db: db,
	}
}

func (r *todoRepository) GetAll(id uint) (todos []domain.Todo) {
	connection := r.db.GetConnection()

	if id == 0 {
		connection.Find(&todos)
	} else {
		connection.Find(&todos, "activity_group_id = ?", id)
	}
	return todos
}

func (r *todoRepository) GetByID(id uint) (todo domain.Todo) {
	connection := r.db.GetConnection()
	connection.First(&todo, "id = ?", id)
	return todo
}

func (r *todoRepository) Create(todo *domain.Todo) error {
	connection := r.db.GetConnection()
	todo.Priority = "very-high"
	result := connection.Create(&todo)
	return result.Error
}

func (r *todoRepository) Update(id uint, todo *domain.Todo) (domain.Todo, error) {
	var old *domain.Todo

	connection := r.db.GetConnection()
	connection.First(&old, "id = ?", id)
	if old.Id == 0 {
		return domain.Todo{}, errors.New("data not found")
	}

	newTodo := map[string]interface{}{
		"Title":    todo.Title,
		"Priority": todo.Priority,
		"IsActive": todo.IsActive,
	}

	result := connection.Model(&old).Updates(&newTodo)
	return *old, result.Error
}

func (r *todoRepository) Delete(id uint) error {
	connection := r.db.GetConnection()
	var old *domain.Todo

	connection.First(&old, "id = ?", id)
	if old.Id == 0 {
		return errors.New("data not found")
	}

	delete := connection.Delete(&old)
	if delete.Error != nil {
		return delete.Error
	}
	return nil
}
