package routes

import (
	"fmt"
	"net/http"
	"plashoes-server/models"

	"github.com/gin-gonic/gin"
)

func submitReview(context *gin.Context){
	var submittedReview models.Review

	err := context.ShouldBindJSON(&submittedReview)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data!", "error": err})
		return
	}

	err = submittedReview.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save review", "error": err})
		return
	}

	newAverageRating,err := submittedReview.GetNewAverageProductRating()

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to get new average rating"})
		return
	}

	err = models.UpdateRating(submittedReview.ProductID, newAverageRating)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to update product rating"})
		return
	}

	err = models.UpdateReviewed(submittedReview.OrderID)

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to update product review"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Product reviewed successfully."})
}