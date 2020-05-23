package domain

type Users []User

type User struct {
	ID     int	`validate:required`
	Amount int
}
