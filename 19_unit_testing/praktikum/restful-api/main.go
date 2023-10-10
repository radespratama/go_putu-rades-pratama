package main

import (
	"fmt"
	"restful-api/src/configs"
	"restful-api/src/database"
	"restful-api/src/routes"
)

func init() {
	// database.Demigrate()
	database.Migrate()
}

func main() {
	api := routes.New()
	serverPort := configs.GetServerConfig().Port
	api.Logger.Fatal(api.Start(fmt.Sprintf(":%s", serverPort)))
}
