package routes

import (
	"fmt"
	"math/rand"
	"net/http"
	"plashoes-server/models"
	"plashoes-server/utils"
	"strconv"

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

func sendOTP(context *gin.Context){
	var requestBody models.OTP

	err := context.ShouldBindJSON(&requestBody)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data!", "error": err})
		return
	}

	userExists, err := models.CheckUserExists(requestBody.Email)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error sending OTP"})
		return
	}

	if userExists {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Email already registered"})
		return
	}

	var randomNumber int
	for randomNumber < 1000 {
		randomNumber = rand.Intn(9999)
	} 
	
	context.Set("OTP", randomNumber)
	
	mailSubject := "Verify mail"
	mailBody := "Use the code " + strconv.Itoa(randomNumber) + " to verify email"

	err = utils.SendMail(requestBody.Email, mailSubject, mailBody)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not send OTP"})
		return
	}

	authToken, err := utils.GenerateToken(requestBody.Email)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not generate auth token", "error": err})
		return
	}

	context.JSON(http.StatusAccepted, gin.H{"authToken": authToken})
}

func verifyOTP(context *gin.Context){
	OTP := context.GetInt("OTP")
	var clientOTP models.ClientOTP

	err := context.ShouldBindJSON(&clientOTP)

	fmt.Println(OTP)
	fmt.Println(clientOTP)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data!", "error": err})
		return
	}

	if OTP != clientOTP.OTP {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Incorrect OTP"})
		return
	}
}

func login(context *gin.Context){
	var loginDetails models.LoginDetails
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

func checkSession(context *gin.Context){
	context.JSON(http.StatusAccepted, gin.H{"message":"Session active"})
}