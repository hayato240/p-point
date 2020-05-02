package main

import (
	"github.com/p-point/infrastructure"
)

func main() {
	infrastructure.Router.Run(":8080")
}
