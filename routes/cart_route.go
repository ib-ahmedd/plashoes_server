package routes

import (
	"fmt"
	"net/http"
	"plashoes-server/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getCartProducts(context *gin.Context) {
	userId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse user id."})
		return
	}

	cartItems,err := models.GetCartItems(userId)

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get cart items.", "error": err})
		return
	}

	context.JSON(http.StatusOK, cartItems)
}

func addCartProduct(context *gin.Context){
	var requestItem models.CartRequest

	err := context.ShouldBindJSON(&requestItem)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data!", "error": err})
		return
	}

	err = requestItem.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could add item to cart", "error": err})
		return
	}
}