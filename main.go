package main

import (
	"github.com/go-basic-api/routes"
)

func main() {
	router := routes.SetupRouter()

	router.Run()
}
