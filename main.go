package main

import (
	"bappa.com/rest/db"
	"bappa.com/rest/routes"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
)

func main() {

	db.InitDB()                // first start the database
	var server = gin.Default() //start the server

	//register all routes
	routes.RegisterRoute(server)


	color.Cyan("🔋 🚀 Server running at http://localhost:8080")
	server.Run(":8080")

}


