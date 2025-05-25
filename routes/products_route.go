package routes

import (
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
	query := "SELECT * FROM PRODUCTS ORDER BY sale ASC LIMIT 6"
	events, err := models.GetProductsFromDB(query)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get best sellers.", "err": err})
		return
	}
	context.JSON(http.StatusOK, events)
}

func getNewArrivals (context *gin.Context) {
	query := "SELECT * FROM PRODUCTS ORDER BY date_arrived ASC LIMIT 9"
	events, err := models.GetProductsFromDB(query)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get best sellers.", "err": err})
		return
	}
	context.JSON(http.StatusOK, events)
}