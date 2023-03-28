package usecase

import "github.com/mjawa20/todo-list-go.git/domain"

type activityUsecase struct {
	repository domain.ActivityRepository
}

func NewActivityUsecase(repository domain.ActivityRepository) domain.ActivityUseCase {
	return &activityUsecase{
		repository: repository,
	}
}

func (u *activityUsecase) GetAll() []domain.Activity {
	return u.repository.GetAll()
}

func (u *activityUsecase) GetByID(id uint) domain.Activity {
	return u.repository.GetByID(id)
}

func (u *activityUsecase) Create(activity *domain.Activity) error {
	return u.repository.Create(activity)
}

func (u *activityUsecase) Update(activity *domain.Activity) error {
	return u.repository.Update(activity)
}

func (u *activityUsecase) Delete(id uint) error {
	return u.repository.Delete(id)
}
