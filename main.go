package main

import (
	"app/component/database"
	"app/routes"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	fmt.Println("Main Application Starts")
	dbConnect := database.ConnectDatabase()
	r := routes.SetupRouter(dbConnect)

	r.Run(":8080")

	signChan := make(chan os.Signal, 1)
	signal.Notify(signChan, os.Interrupt, syscall.SIGTERM)

	<-signChan
	log.Println("Shutting down")
}
