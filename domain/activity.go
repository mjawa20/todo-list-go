package domain

type Activity struct {
	Id    int64  `gorm:"primaryKey;not null" json:"id"`
	Title string `gorm:"type:varchar" json:"title" validate:"required"`
	Email string `gorm:"type:varchar" json:"email"`
	Model
}

type ActivityUseCase interface {
	GetAll() []Activity
	GetByID(id uint) Activity
	Create(activity *Activity) error
	Update(id uint, activity *Activity) (Activity, error)
	Delete(id uint) error
}

type ActivityRepository interface {
	GetAll() []Activity
	GetByID(id uint) Activity
	Create(activity *Activity) error
	Update(id uint, activity *Activity) (Activity, error)
	Delete(id uint) error
}
