package routes

import (
	"plashoes-server/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/best-sellers", getBestSellers)
	server.GET("/new-arrivals", getNewArrivals)
	server.GET("/sale", getSaleProducts)
	server.GET("/product/:id", getSingleProduct)
	server.GET("/product-page/:page", getProducts)
	server.POST("/filter-sort", filterSort)
	server.POST("/otp", sendOTP)
	server.POST("/login", login)
	
	authenticate := server.Group("/")
	authenticate.Use(middlewares.Authenticate)
	authenticate.POST("/verify-otp", verifyOTP)
	authenticate.POST("/register", registerUser)
	authenticate.GET("/cart/:id", getCartProducts)
}