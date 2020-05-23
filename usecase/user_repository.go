package usecase

import "github.com/p-point/domain"

type UserRepository interface {
	Add(domain.User) (int, error)
	FindById(int) (domain.User, error)
	PointUp(domain.User) (int, error)
}
