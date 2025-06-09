package controllers

import (
	"net/http"
	"strconv"

	"bappa.com/rest/models"
	"bappa.com/rest/services"
	"github.com/gin-gonic/gin"
)

func RegisterForEvent(context *gin.Context)  {
	userId :=context.GetInt64("userId")
	eventId,err := strconv.ParseInt(context.Param("id"),10,64)

	if err !=nil {
		context.JSON(http.StatusBadRequest,gin.H{"message":"Could not parse the event id."})
		return
	}

	event, err := models.GetEventById(eventId)
	if err !=nil {
		context.JSON(http.StatusInternalServerError,gin.H{"message":"Could not fetch the event"})
		return
	}

	err = services.RegisterEvent(userId, *event)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register for the event"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Successfully registered for the event"})
	
}

func CancelForEvent(context *gin.Context)  {
	
}