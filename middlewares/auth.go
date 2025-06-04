package middlewares

import (
	"fmt"
	"net/http"
	"plashoes-server/utils"

	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context){
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
		return
	}

	err := utils.VerifyToken(token)
	if err != nil {
		fmt.Println(err)
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
		return
	}
	context.Next()
}