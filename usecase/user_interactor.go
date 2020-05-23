package usecase

import (
	"github.com/p-point/domain"
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
	return user, nil
}

func (interactor *UserInteractor) Show(id int) (user domain.User, err error) {
	user, err = interactor.UserRepository.FindById(id)
	return user, nil
}

func (interactor *UserInteractor) Update(u domain.User) (user domain.User, err error) {
	identify, err := interactor.UserRepository.Update(u)
	if err != nil {
		return user, err
	}

	user, err = interactor.UserRepository.FindById(identify)
	return user, nil
}
