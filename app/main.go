package main

import (
	"MyGram/infrastructure"
	_ "github.com/lib/pq"
)

func main() {
	err := infrastructure.Database.DBInit()
	if err != nil {
		panic(err)
	}
	//_, err = infrastructure.Route.RouterInit()
	//if err != nil {
	//	panic(err)
	//}

}
