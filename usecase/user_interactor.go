package usecase

import (
	"errors"
	"github.com/p-point/domain"
)

type UserInteractor struct {
	UserRepository UserRepository
}

func (interactor *UserInteractor) Add(u domain.User) (user domain.User, err error) {
	_, err = interactor.UserRepository.Add(u)
	if err != nil {
		return  domain.User{}, err
	}
	// TODO: implement find_by(user_id) and return domain.User
	return domain.User{}, errors.New("unimplement")
}