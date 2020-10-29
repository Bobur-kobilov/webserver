package main

import (
	DB "github.com/webserver/persistence"
	router "github.com/webserver/routes"
)

func main() {
	DBInstance := DB.InitDB()
	router.Router(DBInstance)
}
