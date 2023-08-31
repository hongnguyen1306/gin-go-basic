package main

import (
	"app/routes"
	"fmt"
)

func main() {
	fmt.Println("Main Application Starts")
	r := routes.SetupRouter()
	r.Run(":8080")
}
