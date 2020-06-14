package database

import (
	"fmt"
	"log"

	"github.com/hayato240/p-point/domain"
)

type UserRepository struct {
	SqlHandler
}

type TrashScanner struct{}

func (TrashScanner) Scan(interface{}) error {
	return nil
}

func (repo *UserRepository) Add(u domain.User) (id int, err error) {
	result, err := repo.Execute("Insert INTO users (amount) VALUES (?)", u.Amount)
	if err != nil {
		return
	}
	id64, err := result.LastInsertId()
	fmt.Printf("errは%v", err)
	fmt.Printf("id64は%v", id64)
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
		log.Fatal(err)
		return user, err
	}

	var id int
	var amount int
	row.Next()
	if err = row.Scan(
		&id,
		&amount,
		TrashScanner{},
	); err != nil {
		log.Fatal(err)
		return user, err
	}

	user.ID = id
	user.Amount = amount
	return user, nil
}

//Points :adds points to User.
func (repo *UserRepository) Points(u domain.User) (id int, err error) {
	user, err := repo.FindById(u.ID)
	var newAmount int
	newAmount = int(user.Amount) + u.Amount

	_, err = repo.Execute("UPDATE users SET amount = ? WHERE id = ?", newAmount, user.ID)
	if err != nil {
		return
	}

	id = int(user.ID)

	return
}
