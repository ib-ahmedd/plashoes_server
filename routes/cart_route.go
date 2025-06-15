package routes

import (
	"fmt"
	"net/http"
	"plashoes-server/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getCartitems(context *gin.Context) {
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

func addCartitem(context *gin.Context){
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

func updateCartitem(context *gin.Context){

	cartItemID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	var requestItem models.CartRequest
	
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data!", "error": err})
		return
	}
	err = context.ShouldBindJSON(&requestItem)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data!", "error": err})
		return
	}


	err = requestItem.Update(cartItemID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could update item in cart", "error": err})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Cart updated successfully!"})
}

func deleteCartitem(context *gin.Context){
	itemID, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse item id."})
		return
	}

	err = models.DeleteCartItem(itemID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to delete cart item.", "error": err})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Item deleted successfully."})
}

func emptyCart(context *gin.Context){
	userID, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse user id."})
		return
	}

	err = models.EmptyCart(userID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not empty cart."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Cart emptied successfully."})
}