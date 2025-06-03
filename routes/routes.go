package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/best-sellers", getBestSellers)
	server.GET("/new-arrivals", getNewArrivals)
	server.GET("/sale", getSaleProducts)
	server.GET("/product/:id", getSingleProduct)
	server.GET("/product-page/:page", getProducts)
	server.POST("/filter-sort", filterSort)
}