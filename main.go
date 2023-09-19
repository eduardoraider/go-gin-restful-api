package main

import (
	"github.com/eduardoraider/gin-api-rest/database"
	"github.com/eduardoraider/gin-api-rest/routes"
)

func main() {
	database.ConnectDatabase()
	routes.HandleRequests()
}
