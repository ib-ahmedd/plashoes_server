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
	server.POST("/forgot-password", sendForgotOTP)
	server.POST("/login", login)
	
	authenticate := server.Group("/")
	authenticate.Use(middlewares.Authenticate)
	authenticate.GET("/check-session", checkSession)
	authenticate.POST("/verify-otp", verifyOTP)
	authenticate.POST("/register", registerUser)
	authenticate.PATCH("/reset-password", resetPassword)
	authenticate.GET("/cart/:id", getCartitems)
	authenticate.POST("/add-cart", addCartitem)
	authenticate.PATCH("cart-update", updateItemQuantity)
	authenticate.DELETE("/cart-delete/:id", deleteCartitem)
	authenticate.DELETE("/empty-cart/:id", emptyCart)
	authenticate.POST("/order", orderItems)
	authenticate.GET("/orders/:id", getOrders)
	authenticate.GET("/pending-reviews/:id", getPendingReviews)
	authenticate.GET("/order-details/:id", getOrderDetails)
	authenticate.GET("/review/:id", getReviewItem)
	authenticate.POST("/submit-review", submitReview)
}