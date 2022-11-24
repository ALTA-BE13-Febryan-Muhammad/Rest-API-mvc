package main

import (
	"api/mvc/config"
	"api/mvc/routes"
)

func main() {
	config.InitDB()
	config.InitialMigration()

	e := routes.New()

	e.Logger.Fatal(e.Start(":8080"))
}
