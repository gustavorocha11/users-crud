package main

import (
	"github.com/gustavorocha11/users-crud/database"
	"github.com/gustavorocha11/users-crud/server"
)

func main() {
	database.StartDB()
	s := server.NewServer()

	s.Run()
}
