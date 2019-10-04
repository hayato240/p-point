package database

import "github.com/p-point/domain"

type UserRepository struct {
	SqlHandler
}

func (repo *UserRepository) Add(u domain.User) (id int, err error) {
	result, err := repo.Execute("Insert INTO users (amount) VALUES (?)", u.Amount)
	if err != nil{
		return
	}
	id64, err := result.LastInsertId()
	if err != nil {
		return
	}
	id = int(id64)
	return
}

func (repo *UserRepository) FindById(identifier int) (domain.User, error) {
	row, err := repo.Query("SELECT * FROM users WHERE id = ?", identifier)
	user := domain.User{}
	defer row.Close()
	if err != nil {
		return user, err
	}
	var id int
	var amount int
	row.Next()
	if err = row.Scan(&id, &amount); err != nil {
		return user, err
	}
	user.ID = id
	user.Amount = amount
	return user, nil
}