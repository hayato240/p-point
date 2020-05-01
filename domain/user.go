package domain

import {
	"gopkg.in/go-playground/validator.v9"
}

type Users []User

type User struct {
	ID     int	`validate:required`
	Amount int
}
