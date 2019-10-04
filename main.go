package main

import (
	"github.com/p-point/infrastructure"
)

func main() {
	infrastructure.Migrate()
	infrastructure.Router.Run(":8080")
}
