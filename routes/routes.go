package routes

// import "github.com/gin-gonic/gin"

import (
	"bappa.com/rest/controllers"
	"bappa.com/rest/middlewares"
	"github.com/gin-gonic/gin"
)

func  RegisterRoute(server *gin.Engine)  {

	//root: http://localhost:8080
	//event routes


	// Best Practice 
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)

	// events 
	server.GET("/events", controllers.GetEvents)
	server.GET("/events/:id",controllers.GetEvent) 
	authenticated.POST("/events", controllers.CreateEvent)
	authenticated.PUT("/events/:id", controllers.UpdateEvent)
	authenticated.DELETE("/events/:id", controllers.DeleteEvent)

		//users
	server.POST("/signup",controllers.Signup)
	server.POST("/login",controllers.Login)


	//todo:better way
	// server.POST("/events",middlewares.Authenticate, controllers.CreateEvent)
	// server.GET("/events", controllers.GetEvents)
	// server.GET("/events/:id",controllers.GetEvent) 
	// server.PUT("/events/:id", controllers.UpdateEvent)
	// server.DELETE("/events/:id", controllers.DeleteEvent)






	//todo: Old way
	// server.GET("/events", getEvents)
	// server.GET("/events/:id",getEvent) 
	// server.POST("/events", createEvent)
	// server.PUT("/events/:id", updateEvent)
	// server.DELETE("/events/:id", deleteEvent)
}