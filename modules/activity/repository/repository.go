package repository

import (
	"errors"

	"github.com/mjawa20/todo-list-go.git/domain"
)

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

func (r *activityRepository) Update(id uint, activity *domain.Activity) (domain.Activity, error) {
	var old *domain.Activity

	connection := r.db.GetConnection()
	connection.First(&old, "id = ?", id)
	if old.Id == 0 {
		return domain.Activity{}, errors.New("data not found")
	}

	result := connection.Model(&old).Updates(&activity)
	return *old, result.Error
}

func (r *activityRepository) Delete(id uint) error {
	connection := r.db.GetConnection()
	var old *domain.Activity

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
