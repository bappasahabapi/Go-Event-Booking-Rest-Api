package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"bappa.com/rest/models"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not fetch events",
		})
		return // other wise below code will not executed
	}

	// Convert events to pretty JSON
	jsonOutput, _ := json.MarshalIndent(events, "", "  ")
	color.Green("📦 Retrieved Events:\n%s", string(jsonOutput))

	context.JSON(http.StatusOK, events)
}

func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse event id",
		})
		return // other wise below code will not executed
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not fetch event",
		})
		return
	}

	// Convert events to pretty JSON
	jsonOutput, _ := json.MarshalIndent(event, "", "  ")
	color.Green("📦 Retrieved Events:\n%s", string(jsonOutput))

	context.JSON(http.StatusOK, event)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse the request data",
		})
		return
	}

	event.ID = 1    //dummy value
	event.UserID = 1 //dummy value

	//save event
	// err =models.Save(event)  // if use func Save(e Event){...})
	err = event.Save() // if use func(e Event) Save(){...}  // receiver function

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not create event",
		})
		return // other wise below code will not executed
	}

	//also send a response for success case
	context.JSON(http.StatusCreated, gin.H{
		"message": "✅ Event Created Successfully",
		"event":   event,
	})
}

func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)


	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse event id",
		})
		return // other wise below code will not executed
	}

	_, err= models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not fetch the  event",
		})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse the request data",
		})
		return
	}

	updatedEvent.ID=eventId
	err=updatedEvent.Update()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not update event",
		})
		return // other wise below code will not executed
	}
	// log.Printf("🛠 Updating event ID %d with: %+v", eventId, updatedEvent)

	// Convert events to pretty JSON
	jsonOutput, _ := json.MarshalIndent(updatedEvent, "", "  ")
	color.Green("📦 Retrieved Events:\n%s", string(jsonOutput))


	//also send a response for success case
	context.JSON(http.StatusOK, gin.H{
		"message": "✅ Event Updated Successfully",
		"event":   updatedEvent,
	})
}

func deleteEvent(context *gin.Context){
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse event id",
		})
		return // other wise below code will not executed
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not fetch the  event",
		})
		return
	}

	err=event.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not delte the event",
		})
		return // other wise below code will not executed
	}

	//also send a response for success case
	context.JSON(http.StatusOK, gin.H{
		"message": "✅ Event Deleted Successfully",
	})
}
