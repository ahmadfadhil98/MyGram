package main

import (
	"MyGram/database"
	"MyGram/infrastructure"
	"errors"
	_ "github.com/lib/pq"
)

func main() {
	err := database.Database.DBInit()
	if err != nil {
		panic(errors.New("Database connection failed"))
	}
	_, err = infrastructure.Route.RouterInit()
	if err != nil {
		panic(errors.New("Router initialization failed"))
	}

}
