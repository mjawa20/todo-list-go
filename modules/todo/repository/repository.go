package repository

import (
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

	connection.Find(&todos, "activity_group_id = ?", id)
	return todos
}

func (r *todoRepository) GetByID(id uint) (todo domain.Todo) {
	connection := r.db.GetConnection()
	connection.First(&todo, "id = ?", id)
	return todo
}

func (r *todoRepository) Create(todo *domain.Todo) error {
	connection := r.db.GetConnection()
	result := connection.Create(&todo)
	return result.Error
}

func (r *todoRepository) Update(todo *domain.Todo) error {
	var old *domain.Todo

	connection := r.db.GetConnection()
	connection.First(&old, "id = ?", todo.Id)

	old = todo

	result := connection.Save(&old)
	return result.Error
}

func (r *todoRepository) Delete(id uint) error {
	connnection := r.db.GetConnection()
	delete := connnection.Delete(&domain.Todo{Id: int64(id)})
	if delete.Error != nil {
		return delete.Error
	}
	return nil
}
