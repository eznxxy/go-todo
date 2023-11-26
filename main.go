package main

import (
	"github.com/eznxxy/go-todo/database"
	"github.com/eznxxy/go-todo/routes"
)

func main() {
	database.InitDb()

	e := routes.Init()

	e.Logger.Fatal(e.Start(":8080"))
}
