package routes

import (
	"fmt"
	"net/http"
	"plashoes-server/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getSingleProduct(context *gin.Context){
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse product id."})
		return
	}

	result, err := models.GetSingleProduct(eventId)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not get product."})
		return
	}

	context.JSON(http.StatusOK, result)
}

func getBestSellers(context *gin.Context){
	query := "SELECT * FROM PRODUCTS ORDER BY sale ASC LIMIT 6"
	events, err := models.GetProductsFromDB(query)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get best sellers.", "err": err})
		return
	}
	context.JSON(http.StatusOK, events)
}

func getNewArrivals(context *gin.Context) {
	query := "SELECT * FROM PRODUCTS ORDER BY date_arrived ASC LIMIT 9"
	events, err := models.GetProductsFromDB(query)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get best sellers.", "err": err})
		return
	}
	context.JSON(http.StatusOK, events)
}

func getSaleProducts(context *gin.Context){
	query := "SELECT * FROM PRODUCTS WHERE sale = true LIMIT 6"
	events, err := models.GetProductsFromDB(query)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get best sellers.", "err": err})
		return
	}
	context.JSON(http.StatusOK, events)
}

func getProducts(context *gin.Context) {
	page := context.Param("page")
	if page == "Shop" {
		response, err := models.GetAllProducts()
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, "Could not fetch data!")
		return
	}

	context.JSON(http.StatusOK, response)
	return
	}
	response, err := models.GetProductPage(page)
	if err != nil {
		context.JSON(http.StatusInternalServerError, "Could not fetch data!")
		return
	}

	context.JSON(http.StatusOK, response)
}

func filterSort(context *gin.Context){
	var requestBody models.FilterSort
	err := context.ShouldBindJSON(&requestBody)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}
	products, err := models.FilterSortProducts(requestBody)

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not get products."})
		return
	}

	// fmt.Println(products)
	context.JSON(http.StatusOK, products)
}