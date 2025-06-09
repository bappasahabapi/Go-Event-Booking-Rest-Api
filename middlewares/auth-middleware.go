package middlewares

import (
	"net/http"

	"bappa.com/rest/utils"
	"github.com/gin-gonic/gin"
)


func Authenticate(context *gin.Context)  {
	

	// ----- token varification part 

	 token := context.Request.Header.Get("Authorization")

	 if token ==""{
		context.AbortWithStatusJSON(http.StatusUnauthorized,gin.H{"message":"Not Authorized"})
		return // so that below code is not excuated
	 }

	 //if token is failed to varify
	userId,err:=utils.VarifyToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Not Authorized",
		})
		return
	}

	context.Set("userId",userId)
	context.Next() // next handler will work if code comes to this line without error

	// ----- 
}

