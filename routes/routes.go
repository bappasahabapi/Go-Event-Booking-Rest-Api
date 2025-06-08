package routes

// import "github.com/gin-gonic/gin"

import (
	"bappa.com/rest/controllers"
	"github.com/gin-gonic/gin"
)

func  RegisterRoute(server *gin.Engine)  {

	//root: http://localhost:8080
	//event routes
	server.GET("/events", controllers.GetEvents)
	server.GET("/events/:id",controllers.GetEvent) 
	server.POST("/events", controllers.CreateEvent)
	server.PUT("/events/:id", controllers.UpdateEvent)
	server.DELETE("/events/:id", controllers.DeleteEvent)

	//users
	server.POST("/signup",controllers.Signup)

	// server.GET("/events", getEvents)
	// server.GET("/events/:id",getEvent) 
	// server.POST("/events", createEvent)
	// server.PUT("/events/:id", updateEvent)
	// server.DELETE("/events/:id", deleteEvent)
}