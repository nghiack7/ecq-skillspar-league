package main

import (
	"github.com/nghiack7/ecq-skillspar-league/controllers"
	"github.com/nghiack7/ecq-skillspar-league/pkg/models"
)

func main() {
	// init database
	db, err := models.InitDatabase()
	if err != nil {
		panic(err)
	}
	// init server mux or gin
	server := controllers.InitServer(db)
	server.StartServer()
	// run server

	// graceful shutdowm
}
