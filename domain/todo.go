package domain

type Todos struct {
	Id              int64  `gorm:"primaryKey" json:"id"`
	ActivityGroupId int64  `gorm:"type:int" json:"activity_group_id"`
	Title           string `gorm:"type:varchar(255)" json:"title"`
	IsActive        bool   `json:"is_active"`
	Priority        string `gorm:"type:varchar(20)" json:"priority"`
	Model
}

type TodoUseCase interface {
	GetAll(id uint) []Todos
	GetByID(id uint) Todos
	Create(activity *Todos) error
	Update(id uint, activity *Todos) (Todos, error)
	Delete(id uint) error
}

type TodoRepository interface {
	GetAll(id uint) []Todos
	GetByID(id uint) Todos
	Create(activity *Todos) error
	Update(id uint, activity *Todos) (Todos, error)
	Delete(id uint) error
}
