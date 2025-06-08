package routes

// import "github.com/gin-gonic/gin"

import (
	"bappa.com/rest/controllers"
	"github.com/gin-gonic/gin"
)

func  RegisterRoute(server *gin.Engine)  {

	//root: http://localhost:8080
	//event routes
	server.GET("/events", controller.GetEvents)
	server.GET("/events/:id",controller.GetEvent) 
	server.POST("/events", controller.CreateEvent)
	server.PUT("/events/:id", controller.UpdateEvent)
	server.DELETE("/events/:id", controller.DeleteEvent)

	// server.GET("/events", getEvents)
	// server.GET("/events/:id",getEvent) 
	// server.POST("/events", createEvent)
	// server.PUT("/events/:id", updateEvent)
	// server.DELETE("/events/:id", deleteEvent)
}