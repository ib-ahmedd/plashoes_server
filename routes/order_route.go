package routes

import (
	"fmt"
	"net/http"
	"plashoes-server/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func orderItems(context *gin.Context) {
	var orderRequest models.OrderRequest
	err := context.ShouldBindJSON(&orderRequest)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data!", "error": err})
		return
	}

	for index,item := range orderRequest.OrderItems {
		item.UserID = orderRequest.UserID
		item.DateOrdered = time.Now().String()
		item.DateDelivered = time.Now().String()
		if index % 2 == 0 {
			item.OrderStatus = "Delivered"
		}else{
			item.OrderStatus = "Processing"
		}

		err = item.Save()

		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save ordered item."})
			return
		}
	}

	context.JSON(http.StatusOK, gin.H{"message": "Items ordered successfully."})
}

func getOrders(context *gin.Context){
	userID, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse user id."})
		return
	}

	orders,err := models.GetOrders(userID)

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to get orders."})
		return
	}

	context.JSON(http.StatusOK, orders)
}

func getPendingReviews(context *gin.Context){
	userID, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse user id."})
		return
	}

	orders,err := models.GetPendingReviews(userID)

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to get items."})
		return
	}

	context.JSON(http.StatusOK, orders)
}

func getOrderDetails(context *gin.Context){
	userID, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse order id."})
		return
	}

	orderDetails,err := models.GetOrderDetails(userID)

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to get order details."})
		return
	}

	context.JSON(http.StatusOK, orderDetails)
}