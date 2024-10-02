package main

import (
	"github.com/olherasymchuk/bookovyna/data"
	"github.com/olherasymchuk/bookovyna/handlers"
)

func main() {

	data.ConnectDB()
	data.LoadMock()

	r := handlers.SetupRouter()
	r.Run("localhost:8080")
}
