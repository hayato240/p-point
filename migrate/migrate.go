package main

import (
	"github.com/hayato240/p-point/interfaces/database"
	"github.com/hayato240/p-point/domain"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db := database.Connection()
	defer db.Close()

	db.AutoMigrate(&domain.User{})
}
