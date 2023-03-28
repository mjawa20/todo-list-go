package domain

type Todo struct {
	Id              int64  `gorm:"primaryKey" json:"id"`
	ActivityGroupId int64  `gorm:"type:int" json:"activity_group_id"`
	Title           string `gorm:"type:varchar" json:"title"`
	IsActive        bool   `gorm:"type:boolean" json:"is_active"`
	Priority        string `gorm:"type:varchar(20)" json:"priority"`
	Model
}

type TodoUseCase interface {
	GetAll() []Todo
	GetByID(id uint) Todo
	Create(todo *Todo) error
	Update(todo *Todo) error
	Delete(id uint) error
}

type TodoRepository interface {
	GetAll() []Todo
	GetByID(id uint) Todo
	Create(todo *Todo) error
	Update(todo *Todo) error
	Delete(id uint) error
}