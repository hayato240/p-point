package main

import (
	"github.com/hayato240/p-point/infrastructure"
)

func main() {
	infrastructure.Router.Run(":8080")
}
