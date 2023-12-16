package main

import (
	"Assignment-3/controller"
	"Assignment-3/database"
	"Assignment-3/routes"
)

func main() {

	database.StartDB()

	go controller.UpdateSensorData()

	r := routes.SetupRoutes()
	r.Run(":8080")
}
