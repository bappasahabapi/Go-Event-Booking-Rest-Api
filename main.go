package main

import (
	"fmt"
	"net/http"

	"bappa.com/rest/models"
	"github.com/gin-gonic/gin"
)

func main() {

	var server = gin.Default()

	server.GET("/events",getEvents)
	server.POST("/events",createEvent)

	fmt.Println("🚀 Server running at http://localhost:8080")
    server.Run(":8080")

}

func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
	context.JSON(http.StatusOK,events)
}

func createEvent(context *gin.Context){
	var event models.Event
	err :=context.ShouldBindJSON(&event)

	if err !=nil {
		context.JSON(http.StatusBadRequest,gin.H{
			"message": "Could not parse the request data",
		})
		return
	}

	event.ID=1 //dummy value
	event.UserID = 1 //dummy value

	//save event
	event.SaveEvent()

	//also send a response for success case
	context.JSON(http.StatusCreated,gin.H{
		"message":"✅ Event Created Successfully",
		"event":event,
	})
}
