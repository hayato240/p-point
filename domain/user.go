package domain

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	amount      int       `json:amount`
}
