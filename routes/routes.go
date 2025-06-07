package routes

import "github.com/gin-gonic/gin"

func  RegisterRoute(server *gin.Engine)  {

	//root: http://localhost:8080
	//event routes
	server.GET("/events", getEvents)
	server.GET("/events/:id",getEvent) 
	server.POST("/events", createEvent)
	server.PUT("/events/:id", updateEvent)
	server.DELETE("/events/:id", deleteEvent)
}