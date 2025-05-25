package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/best-sellers", getBestSellers)
	server.GET("/new-arrivals", getNewArrivals)
}