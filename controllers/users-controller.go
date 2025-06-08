package controllers

import (
	"net/http"

	"bappa.com/rest/models"
	"bappa.com/rest/services"
	"github.com/gin-gonic/gin"
)

func Signup(context *gin.Context){
	var user models.User
	err :=context.ShouldBindJSON(&user)

	println(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse the request data",
		})
		return
	}

	err =services.SaveUser(&user)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not create user",
		})
		return // other wise below code will not executed
	}

	//also send a response for success case
	context.JSON(http.StatusCreated, gin.H{
		"message": "✅ User Created Successfully",
		"event":   user,
	})
}

func __Signup(context *gin.Context) {
	var user models.User

	if err := context.ShouldBindJSON(&user); err != nil {
		// Log the exact error for debugging
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse the request data",
			"error":   err.Error(), // <- this line is new
		})
		return
	}

	if err := services.SaveUser(&user); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not create user",
			"error":   err.Error(), // <- log DB error too
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "✅ User Created Successfully",
		"user":    user,
	})
}

func Login(context *gin.Context){
	var user models.User
	// println("User",&user)
	err:=context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse the request data",
		})
		return
	}


	//validate
	err =services.ValidateCredentials(&user)
	if err !=nil {
		context.JSON(http.StatusUnauthorized,gin.H{
			"message":"Could not authenticate the user",
			// "message":err.Error(),
		})
		return
	}

	//valid user
	context.JSON(http.StatusCreated, gin.H{
		"message": "✅ User Logged in Successfully",
		"event":   user,
	})

}
