package routes

import (
	"fmt"
	"net/http"
	"plashoes-server/models"

	"github.com/gin-gonic/gin"
)

func getAllProducts(context *gin.Context) {
	events, err := models.GetAllProducts()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get all products.", "err": err})
		return
	}
	context.JSON(http.StatusOK, events)
}

func getBestSellers (context *gin.Context){
	events, err := models.GetBestSellers()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get best sellers.", "err": err})
		return
	}

	fmt.Println(events)
	context.JSON(http.StatusOK, events)
}