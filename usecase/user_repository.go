package usecase

import "github.com/p-point/domain"

type UserRepository interface {
	Add(domain.User) (int, error)
	FindById(int) (domain.User, error)
	Points(domain.User) (int, error)
}
