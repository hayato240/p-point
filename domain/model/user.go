package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	amount      int       `json:age`
}
