package routes

import (
	"fmt"
	"net/http"
	"plashoes-server/models"
	"plashoes-server/utils"

	"github.com/gin-gonic/gin"
)

func registerUser(context *gin.Context) {

	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data!", "error": err})
		return
	}

	newUser,err := user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user.", "error": err})
		return
	}

	accessToken, err := utils.GenerateToken(newUser.Email)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not generate access token.", "error": err})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"userInfo": newUser, "accessToken": accessToken })
}

func login(context *gin.Context){
	var loginDetails models.User
	err := context.ShouldBindJSON(&loginDetails)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data!", "error": err})
		return
	}

	userExists, err := models.CheckUserExists(loginDetails.Email)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error reading database."})
		return
	}

	if !userExists {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Email not registered."})
		return
	}

	user,err := loginDetails.Login()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "User name or password incorect."})
		return
	}

	accessToken, err := utils.GenerateToken(user.Email)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to generate access token"})
		return
	}

	context.JSON(http.StatusAccepted, gin.H{"userInfo": user, "accessToken": accessToken})
}

func resetPassword(context *gin.Context){
	var resetRequest models.User

	err := context.ShouldBindJSON(&resetRequest)

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Unable to parse request data."})
		return
	}

	err = resetRequest.ResetPassword()

	if err != nil{
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error resetting password."})
	}

	err = models.DeleteOTP(resetRequest.Email)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error deleting user OTP from database."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Password reset succesful."})
}

func checkSession(context *gin.Context){
	context.JSON(http.StatusAccepted, gin.H{"message":"Session active"})
}