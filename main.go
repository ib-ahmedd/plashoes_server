package main

import (
	// "fmt"
	"plashoes-server/db"
	// "plashoes-server/models"
	"plashoes-server/routes"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	corsConfig := cors.Config{
		AllowAllOrigins: true,
		AllowMethods: []string {"POST", "GET", "UPDATE", "DELETE", "PUT"},
		AllowHeaders: []string {"Content-Type", "Authorization"},
		ExposeHeaders: []string {"Content-Length"},
		AllowCredentials: true,
		MaxAge: 12 * time.Hour,

	}

	server.Use(cors.New(corsConfig))
	routes.RegisterRoutes(server)
	server.Run(":8080")
	

	// for _,item := range models.ProductsArray {
	// 	err := item.Save()
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		break
	// 	}

	// 	fmt.Println("Save success")
	// }

}