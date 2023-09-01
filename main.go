package main

import (
	"app/component/database"
	"app/routes"
	"fmt"
)

func main() {
	fmt.Println("Main Application Starts")
	dbConnect := database.ConnectDatabase()

	r := routes.SetupRouter(dbConnect)
	r.Run(":8080")
}
