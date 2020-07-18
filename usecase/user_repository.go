package usecase

import (
	"database/sql"

	"github.com/hayato240/p-point/domain"
)

type UserRepository interface {
	Add(domain.User) (int, error)
	FindById(int) (domain.User, error)
	Points(domain.User) (int, error)
	UpdateAmout(*sql.Tx, int, int) error
}
