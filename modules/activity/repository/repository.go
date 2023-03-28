package repository

import "github.com/mjawa20/todo-list-go.git/domain"

type activityRepository struct {
	db domain.DB
}

func NewActivityRepository(db domain.DB) domain.ActivityRepository {
	return &activityRepository{
		db: db,
	}
}

func (r *activityRepository) GetAll() (activities []domain.Activity) {
	connection := r.db.GetConnection()
	connection.Find(&activities)
	return activities
}

func (r *activityRepository) GetByID(id uint) (activity domain.Activity) {
	connection := r.db.GetConnection()
	connection.First(&activity, "id = ?", id)
	return activity
}

func (r *activityRepository) Create(activity *domain.Activity) error {
	connection := r.db.GetConnection()
	result := connection.Create(&activity)
	return result.Error
}

func (r *activityRepository) Update(activity *domain.Activity) error {
	var old *domain.Activity

	connection := r.db.GetConnection()
	connection.First(&old, "id = ?", activity.Id)

	old = activity

	result := connection.Save(&old)
	return result.Error
}

func (r *activityRepository) Delete(id uint) error {
	connnection := r.db.GetConnection()
	delete := connnection.Delete(&domain.Activity{Id: int64(id)})
	if delete.Error != nil {
		return delete.Error
	}
	return nil
}
