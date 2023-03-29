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

func (u *activityUsecase) GetAll() []domain.Activities {
	return u.repository.GetAll()
}

func (u *activityUsecase) GetByID(id uint) domain.Activities {
	return u.repository.GetByID(id)
}

func (u *activityUsecase) Create(activity *domain.Activities) error {
	return u.repository.Create(activity)
}

func (u *activityUsecase) Update(id uint, activity *domain.Activities) (domain.Activities, error) {
	return u.repository.Update(id, activity)
}

func (u *activityUsecase) Delete(id uint) error {
	return u.repository.Delete(id)
}
