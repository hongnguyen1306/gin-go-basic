package main

import (
	"app/db"
	"app/routes"
	"log"
)

func main() {
	db.ConnectDatabase()
	log.Fatal(routes.Routes(":8080"))
}
