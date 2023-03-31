package main

import (
	"todo-app/database"
	"todo-app/routers"
)

func main() {
	database.StartDB()
	r := routers.StartApp()
	r.Run(":3030")
}
