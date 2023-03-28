package domain

type Activity struct {
	Id    int64  `gorm:"primaryKey" json:"id"`
	Title string `gorm:"type:varchar" json:"title"`
	Email string `gorm:"type:varchar" json:"email"`
	Model
}

type ActivityUseCase interface {
	GetAll() []Activity
	GetByID(id uint) Activity
	Create(activity *Activity) error
	Update(activity *Activity) error
	Delete(id uint) error
}

type ActivityRepository interface {
	GetAll() []Activity
	GetByID(id uint) Activity
	Create(activity *Activity) error
	Update(activity *Activity) error
	Delete(id uint) error
}
