package main

import (
	"github.com/go-basic-api/config"
	"github.com/go-basic-api/models"
	"github.com/go-basic-api/routes"
)

func main() {
	conn, err := config.StartDBConn(config.MONGO_URI)
	if err != nil {
		println("Error initializing MongoDB")
		println(err.Error())
	}
	models.S.Connect(conn)

	router := routes.SetupRouter()

	router.Run()
}
