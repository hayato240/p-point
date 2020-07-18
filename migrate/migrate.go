package main

import (
	"github.com/hayato240/p-point/interfaces/database"
	"github.com/hayato240/p-point/domain"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Usage: go run /go/src/github.com/hayato240/p-point/migrate/migrate.go
func main() {
	db := database.Connection()
	defer db.Close()

	db.AutoMigrate(&domain.User{})
	db.AutoMigrate(&domain.PointHistory{})
}
