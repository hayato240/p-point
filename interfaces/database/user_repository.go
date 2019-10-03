package database

import "github.com/p-point/domain"

type UserRepository struct {
	SqlHandler
}

func (repo *UserRepository) Add(u domain.User) (id int, err error) {
	result, err := repo.Execute("Insert INTO users (amount) VALUES (?)", u.Amount,)
	if err != nil{
		return
	}
	id64, err := result.LastInsertId()
	if err != nil {
		return
	}
	id = int(id64)
	return id, nil
}