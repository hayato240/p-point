package database

import (
	"database/sql"
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
	if err != nil {
		return
	}
	id = int(id64)
	return id, nil
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
	); err != nil {
		log.Fatal(err)
		return user, err
	}

	user.ID = id
	user.Amount = amount
	return user, nil
}

func (repo *UserRepository) AddPoints(u domain.User) (id int, err error) {
	err = repo.SqlHandler.Transaction(func(tx *sql.Tx) error {
		user, err := repo.FindById(u.ID)
		if err != nil {
			return err
		}
		var newAmount int
		newAmount = user.Amount + u.Amount
		log.Printf("a:: %#v", newAmount)
		err = repo.UpdateAmount(tx, newAmount, user.ID)
		if err != nil {
			return err
		}
		err = repo.AddPointHistory(tx, user.ID, u.Amount)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Fatal("panic")
		log.Fatal(err)
	}
	return u.ID, nil
}

func (repo *UserRepository) UpdateAmount(tx *sql.Tx, newAmount int, userID int) error {
	_, err := tx.Exec("UPDATE users SET amount = ? WHERE id = ?", newAmount, userID) // TODO(Sho): method化 7/18
	return err
}

func (repo *UserRepository) AddPointHistory(tx *sql.Tx, userID int, addedAmount int) error {
	_, err := tx.Exec("INSERT INTO point_histories (user_id, amount) VALUES (?, ?)", userID, addedAmount) // TODO(Sho): method化 7/18
	return err
}
