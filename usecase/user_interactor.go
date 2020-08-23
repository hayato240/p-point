package usecase

import (
	"github.com/hayato240/p-point/domain"
)

type UserInteractor struct {
	UserRepository UserRepository
}

func (interactor *UserInteractor) Add(u domain.User) (user domain.User, err error) {
	identify, err := interactor.UserRepository.Add(u)
	if err != nil {
		return user, err
	}
	user, err = interactor.UserRepository.FindById(identify)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (interactor *UserInteractor) Show(id int) (user domain.User, err error) {
	user, err = interactor.UserRepository.FindById(id)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (interactor *UserInteractor) AddPoints(u domain.User) (user domain.User, err error) {
	identify, err := interactor.UserRepository.AddPoints(u)
	if err != nil {
		return user, err
	}

	user, err = interactor.UserRepository.FindById(identify)
	if err != nil {
		return user, err
	}
	return user, nil
}
