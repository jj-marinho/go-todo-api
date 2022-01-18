package main

import (
	"github.com/go-basic-api/config"
	"github.com/go-basic-api/routes"
)

func main() {
	config.StartMongoDB(config.MONGO_URI)
	router := routes.SetupRouter()

	router.Run()
}
