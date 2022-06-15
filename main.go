package main

import (
	"echo-rest-postgres/db"
	"echo-rest-postgres/routes"
)

func main() {
	db.Init()
	e := routes.Init()

	e.Logger.Fatal(e.Start(":1234"))

}
