package main

import (
	"app/common"
	"app/component/app_context"
	"app/component/database"
	"app/routes"
	"fmt"
)

func main() {
	fmt.Println("Main Application Starts")
	config := common.NewConfig()
	dbConnect := database.New(config)
	appCtx := app_context.NewAppContext(dbConnect, config)
	r := routes.SetupRouter(appCtx)

	r.Run(":8080")
}
