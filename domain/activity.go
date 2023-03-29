package domain

type Activities struct {
	Id    int64  `gorm:"primaryKey" json:"activity_id"`
	Title string `gorm:"type:varchar(255)" json:"title"`
	Email string `gorm:"type:varchar(255)" json:"email"`
	Model
}

type ActivityUseCase interface {
	GetAll() []Activities
	GetByID(id uint) Activities
	Create(activity *Activities) error
	Update(id uint, activity *Activities) (Activities, error)
	Delete(id uint) error
}

type ActivityRepository interface {
	GetAll() []Activities
	GetByID(id uint) Activities
	Create(activity *Activities) error
	Update(id uint, activity *Activities) (Activities, error)
	Delete(id uint) error
}
