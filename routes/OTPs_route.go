package routes

import (
	"math/rand"
	"net/http"
	"plashoes-server/models"
	"plashoes-server/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func verifyOTP(context *gin.Context) {
	var verificationRequest models.OTP

	err := context.ShouldBindJSON(&verificationRequest)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data!", "error": err})
		return
	}

	emailVerified,err := verificationRequest.OTPCorrect()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message":"Error querying database."})
		return
	}

	if !emailVerified {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Incorrect OTP"})
		return
	}

	context.JSON(http.StatusAccepted, gin.H{"message":"Email verified"})
}

func sendForgotOTP(context *gin.Context){
	var requestBody models.OTP

	err := context.ShouldBindJSON(&requestBody)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data!", "error": err})
		return
	}

	userExists, err := models.CheckUserExists(requestBody.Email)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error querying email"})
		return
	}

	if !userExists {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Email not registered"})
		return
	}

	var randomNumber int
	for randomNumber < 1000 {
		randomNumber = rand.Intn(9999)
	} 

	requestBody.Code = randomNumber

	err = requestBody.Save()
	
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message":"unable to add OTP to database"})
		return
	}
	
	
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

func sendOTP(context *gin.Context){
	var requestBody models.OTP

	err := context.ShouldBindJSON(&requestBody)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data!", "error": err})
		return
	}

	userExists, err := models.CheckUserExists(requestBody.Email)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error querying email"})
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
	
	requestBody.Code = randomNumber

	err = requestBody.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message":"unable to add OTP to database"})
		return
	}
	
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

	context.JSON(http.StatusCreated, gin.H{"authToken": authToken})
}