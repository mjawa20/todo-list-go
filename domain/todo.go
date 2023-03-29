package domain

type Todo struct {
	Id              int64  `gorm:"primaryKey;not null" json:"id"`
	ActivityGroupId int64  `gorm:"type:int" json:"activity_group_id"`
	Title           string `gorm:"type:varchar" json:"title"`
	IsActive        bool   `gorm:"type:boolean" json:"is_active"`
	Priority        string `gorm:"type:varchar(20)" json:"priority"`
	Model
}

type TodoUseCase interface {
	GetAll(id uint) []Todo
	GetByID(id uint) Todo
	Create(activity *Todo) error
	Update(id uint, activity *Todo) (Todo, error)
	Delete(id uint) error
}

type TodoRepository interface {
	GetAll(id uint) []Todo
	GetByID(id uint) Todo
	Create(activity *Todo) error
	Update(id uint, activity *Todo) (Todo, error)
	Delete(id uint) error
}
