package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	var server = gin.Default()

	server.GET("/events",getEvents)

	fmt.Println("🚀 Server running at http://localhost:8080")
    server.Run(":8080")

}

func getEvents(context *gin.Context) {
	context.JSON(
		(http.StatusOK),
		gin.H{
			"message": "Events retrieved successfully",
		});

}
