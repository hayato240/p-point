package usecase

import (
	"database/sql"

	"github.com/hayato240/p-point/domain"
)

type UserRepository interface {
	Add(domain.User) (int, error)
	FindById(int) (domain.User, error)
	AddPoints(domain.User) (int, error)
	UpdateAmount(*sql.Tx, int, int) error
	AddPointHistory(*sql.Tx, int, int) error
}
